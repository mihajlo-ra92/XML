package main

import (
	"Rest/data"
	"Rest/handlers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	//Reading from environment, if not set we will default it to 8080.
	//This allows flexibility in different environments (for eg. when running multiple docker api's and want to override the default port)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	// Initialize context
	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//Initialize the logger we are going to use, with prefix and datetime for every log
	logger := log.New(os.Stdout, "[product-api] ", log.LstdFlags)
	storeLogger := log.New(os.Stdout, "[user-store] ", log.LstdFlags)
	flightLogger := log.New(os.Stdout, "[flight-store] ", log.LstdFlags)
	ticketLogger := log.New(os.Stdout, "[ticket-store] ", log.LstdFlags)

	// NoSQL: Initialize Product Repository store
	store, err := data.NewUserRepo(timeoutContext, storeLogger)
	if err != nil {
		logger.Fatal(err)
	}
	defer store.DisconnectUserRepo(timeoutContext)

	flightStore, err := data.NewFlightRepo(timeoutContext, flightLogger)
	if err != nil {
		logger.Fatal(err)
	}
	defer flightStore.DisconnectFlightRepo(timeoutContext)

	ticketStore, err := data.NewTicketRepo(timeoutContext, ticketLogger)
	if err != nil {
		logger.Fatal(err)
	}
	defer ticketStore.DisconnectTicketRepo(timeoutContext)

	// NoSQL: Checking if the connection was established
	store.Ping()
	flightStore.Ping()
	ticketStore.Ping()

	//Initialize the handler and inject said logger
	usersHandler := handlers.NewUsersHandler(logger, store)
	flightsHandler := handlers.NewFlightsHandler(logger, flightStore)
	ticketsHandler := handlers.NewTicketsHandler(logger, ticketStore, flightStore)
	cors := gorillaHandlers.CORS(
		gorillaHandlers.AllowedHeaders([]string{"Bearer", "Content-Type"}),
		gorillaHandlers.AllowedOrigins([]string{"*"}))

	//Initialize the router and add a middleware for all the requests
	router := mux.NewRouter()
	router.Use(cors)
	router.Use(usersHandler.MiddlewareContentTypeSet)

	//NOTE: User routers
	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/user", usersHandler.GetAllUsers)
	getRouter.Use(usersHandler.MiddlewareAuth)

	loginRouter := router.Methods(http.MethodPost).Subrouter()
	loginRouter.HandleFunc("/login", usersHandler.Login)
	loginRouter.Use(usersHandler.MiddlewareLoginDeserialization)

	initRouter := router.Methods(http.MethodGet).Subrouter()
	initRouter.HandleFunc("/init", usersHandler.InitTestDb)

	initFlightRouter := router.Methods(http.MethodGet).Subrouter()
	initFlightRouter.HandleFunc("/init-flight", flightsHandler.InitTestDb)

	initTicketRouter := router.Methods(http.MethodGet).Subrouter()
	initTicketRouter.HandleFunc("/init-ticket", ticketsHandler.InitTestDb)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/user", usersHandler.PostUser)
	postRouter.Use(usersHandler.MiddlewareUserDeserialization)

	getByUsernameRouter := router.Methods(http.MethodGet).Subrouter()
	getByUsernameRouter.HandleFunc("/user/read-by-username", usersHandler.GetUsersByUsername)

	patchUserRouter := router.Methods(http.MethodPatch).Subrouter()
	patchUserRouter.HandleFunc("/user/{id}", usersHandler.PatchUser)
	patchUserRouter.Use(usersHandler.MiddlewareUserDeserialization)

	deleteUserRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteUserRouter.HandleFunc("/user/{id}", usersHandler.DeleteUser)

	//NOTE: Flight routers
	postFlightRouter := router.Methods(http.MethodPost).Subrouter()
	postFlightRouter.HandleFunc("/flight", flightsHandler.PostFlight)
	postFlightRouter.Use(usersHandler.MiddlewareAuth)
	postFlightRouter.Use(flightsHandler.MiddlewareFlightDeserialization)

	getFlightRouter := router.Methods(http.MethodGet).Subrouter()
	getFlightRouter.HandleFunc("/flight", flightsHandler.GetAllFlights)

	getFlightByIdRouter := router.Methods(http.MethodGet).Subrouter()
	getFlightByIdRouter.HandleFunc("/get-flight-by-id", flightsHandler.GetFlightById)

	patchFlightRouter := router.Methods(http.MethodPatch).Subrouter()
	patchFlightRouter.HandleFunc("/flight/{id}", flightsHandler.PatchFlight)
	patchFlightRouter.Use(flightsHandler.MiddlewareFlightDeserialization)

	deleteFlightRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteFlightRouter.HandleFunc("/flight/{id}", flightsHandler.DeleteFlight)

	getByStartPlaceRouter := router.Methods(http.MethodGet).Subrouter()
	getByStartPlaceRouter.HandleFunc("/flight/search", flightsHandler.SearchFlights)


	//NOTE: Ticket routers
	postTicketRouter := router.Methods(http.MethodPost).Subrouter()
	postTicketRouter.HandleFunc("/ticket", ticketsHandler.PostTicket)
	// postTicketRouter.Use(usersHandler.MiddlewareAuth)
	postTicketRouter.Use(ticketsHandler.MiddlewareTicketDeserialization)

	getTicketRouter := router.Methods(http.MethodGet).Subrouter()
	getTicketRouter.HandleFunc("/tickets", ticketsHandler.GetAllTickets)

	getTicketByIdRouter := router.Methods(http.MethodGet).Subrouter()
	getTicketByIdRouter.HandleFunc("/get-ticket-by-id", ticketsHandler.GetTicketById)

	patchTicketRouter := router.Methods(http.MethodPatch).Subrouter()
	patchTicketRouter.HandleFunc("/ticket/{id}", ticketsHandler.PatchTicket)
	patchTicketRouter.Use(ticketsHandler.MiddlewareTicketDeserialization)

	getTicketsByUserIdRouter := router.Methods(http.MethodGet).Subrouter()
	getTicketsByUserIdRouter.HandleFunc("/get-tickets-by-user-id", ticketsHandler.GetTicketsByUserId)
	getTicketsByUserIdRouter.Use(usersHandler.MiddlewareAuth)

	deleteTicketRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteTicketRouter.HandleFunc("/ticket/{id}", ticketsHandler.DeleteTicket)

	//cors := gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"*"}))

	//Initialize the server
	server := http.Server{
		Addr:         ":" + port,
		Handler:      cors(router),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	logger.Println("Server listening on port", port)
	//Distribute all the connections to goroutines
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt)
	signal.Notify(sigCh, os.Kill)

	sig := <-sigCh
	logger.Println("Received terminate, graceful shutdown", sig)

	//Try to shutdown gracefully
	if server.Shutdown(timeoutContext) != nil {
		logger.Fatal("Cannot gracefully shutdown...")
	}
	logger.Println("Server stopped")
}
