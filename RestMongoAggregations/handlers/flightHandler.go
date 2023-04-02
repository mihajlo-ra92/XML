package handlers

import (
	"Rest/data"
	"context"
	"log"
	"net/http"
	"time"
	"strconv"

	"github.com/gorilla/mux"
)

type FlightsHandler struct {
	logger *log.Logger
	repo   *data.FlightRepo
}

func NewFlightsHandler(l *log.Logger, r *data.FlightRepo) *FlightsHandler {
	return &FlightsHandler{l, r}
}

func (u *FlightsHandler) MiddlewareFlightDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		flight := &data.Flight{}
		err := flight.FromJSON(h.Body)
		u.logger.Println(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			u.logger.Fatal(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProduct{}, flight)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}

func (f *FlightsHandler) SearchFlights(rw http.ResponseWriter, h *http.Request){
	startPlace := h.URL.Query().Get("startPlace")
	endPlace := h.URL.Query().Get("endPlace")
	startDateString := h.URL.Query().Get("startDate")
	endDateString := h.URL.Query().Get("endDate")
	quantityString := h.URL.Query().Get("quantity")
	
	
	layout := "2006-01-02T15:04:05.999Z"
	
	startDate, err1 := time.Parse(layout, startDateString)
	endDate, err2 := time.Parse(layout, endDateString)
	
	quantity, errNum := strconv.ParseInt(quantityString,10,0)
	

	if errNum != nil{
		f.logger.Print("Parsing exception: ", err1)
	} 

	if err1 != nil {
		f.logger.Print("Parsing exception: ", err1)
	}

	if err2 != nil {
		f.logger.Print("Parsing exception: ", err2)
	}
	
	flights,err := f.repo.GetByPlaces(startPlace, endPlace, startDate, endDate, int(quantity))
	if err != nil {
		f.logger.Print("Database exception: ", err)
	}
	if flights == nil{
		return
	}
	err = flights.ToJSON(rw)
	if err != nil{
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		f.logger.Fatal("Unable to convert to json: ", err)
		return
	}
}

func (f *FlightsHandler) InitTestDb(rw http.ResponseWriter, h *http.Request) {
	err := f.repo.DropCollection()
	if err != nil {
		f.logger.Print("Database exception: ", err)
	}
	flight1 := data.Flight{
		Date: time.Unix(1735689600,0),
		EndPlace: "Belgrade",
		StartPlace: "Subotica",
		Capacity: 100,
		Price: 111,
		FreeSeats: 99,
	}
	f.repo.Insert(&flight1)

	flight2 := data.Flight{
		Date: time.Unix(1735689600,0),
		EndPlace: "Belgrade",
		StartPlace: "Novi Sad",
		Capacity: 90,
		Price: 112,
		FreeSeats: 90,
	}
	f.repo.Insert(&flight2)

}

func (f *FlightsHandler) PostFlight(rw http.ResponseWriter, h *http.Request) {
	userType := h.Header.Get("userType")
	f.logger.Println(userType)

	if userType == "admin" {
		flight := h.Context().Value(KeyProduct{}).(*data.Flight)
		f.repo.Insert(flight)
		rw.WriteHeader(http.StatusCreated)
	} else {
		f.logger.Println("Not admin")
		rw.WriteHeader(http.StatusUnauthorized)

	}
}

func (f *FlightsHandler) GetAllFlights(rw http.ResponseWriter, h *http.Request) {
	flights, err := f.repo.GetAll()
	if err != nil {
		f.logger.Print("Database exception: ", err)
	}

	if flights == nil {
		return
	}

	err = flights.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		f.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (f *FlightsHandler) GetFlightById(rw http.ResponseWriter, h *http.Request) {
	id := h.URL.Query().Get("id")

	flights, err := f.repo.GetFlightById(id)
	if err != nil {
		f.logger.Print("Database exception: ", err)
	}

	if flights == nil {
		return
	}

	err = flights.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		f.logger.Fatal("Unable to convert to json: ", err)
		return
	}
}

func (u *FlightsHandler) PatchFlight(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	flight := h.Context().Value(KeyProduct{}).(*data.Flight)

	u.repo.Update(id, flight)
	rw.WriteHeader(http.StatusOK)
}

func (f *FlightsHandler) DeleteFlight(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]

	f.repo.Delete(id)
	rw.WriteHeader(http.StatusNoContent)
}
