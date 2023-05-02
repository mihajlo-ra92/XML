package startup

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/mihajlo-ra92/XML/api_gateway/infrastructure/api"
	cfg "github.com/mihajlo-ra92/XML/api_gateway/startup/config"
	accommodationGw "github.com/mihajlo-ra92/XML/common/proto/accommodation_service"
	authGw "github.com/mihajlo-ra92/XML/common/proto/auth_service"
	catalogueGw "github.com/mihajlo-ra92/XML/common/proto/catalogue_service"
	inventoryGw "github.com/mihajlo-ra92/XML/common/proto/inventory_service"
	orderingGw "github.com/mihajlo-ra92/XML/common/proto/ordering_service"
	shippingGw "github.com/mihajlo-ra92/XML/common/proto/shipping_service"
	userGw "github.com/mihajlo-ra92/XML/common/proto/user_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	config *cfg.Config
	mux    *runtime.ServeMux
}

func NewServer(config *cfg.Config) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}
	server.initHandlers()
	server.initCustomHandlers()
	return server
}

func (server *Server) initHandlers() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	catalogueEmdpoint := fmt.Sprintf("%s:%s", server.config.CatalogueHost, server.config.CataloguePort)
	fmt.Print("catalogueEmdpoint: ")
	fmt.Println(catalogueEmdpoint)
	err := catalogueGw.RegisterCatalogueServiceHandlerFromEndpoint(context.TODO(), server.mux, catalogueEmdpoint, opts)
	if err != nil {
		panic(err)
	}
	orderingEmdpoint := fmt.Sprintf("%s:%s", server.config.OrderingHost, server.config.OrderingPort)
	err = orderingGw.RegisterOrderingServiceHandlerFromEndpoint(context.TODO(), server.mux, orderingEmdpoint, opts)
	if err != nil {
		panic(err)
	}
	shippingEmdpoint := fmt.Sprintf("%s:%s", server.config.ShippingHost, server.config.ShippingPort)
	err = shippingGw.RegisterShippingServiceHandlerFromEndpoint(context.TODO(), server.mux, shippingEmdpoint, opts)
	if err != nil {
		panic(err)
	}
	inventoryEndpoint := fmt.Sprintf("%s:%s", server.config.InventoryHost, server.config.InventoryPort)
	err = inventoryGw.RegisterInventoryServiceHandlerFromEndpoint(context.TODO(), server.mux, inventoryEndpoint, opts)
	if err != nil {
		panic(err)
	}
	
	//NOTE: My endpoints
	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	fmt.Print("userEndpoint: ")
	fmt.Println(userEndpoint)
	err = userGw.RegisterUserServiceHandlerFromEndpoint(context.TODO(), server.mux, userEndpoint, opts)
	if err != nil {
		panic(err)
	}

	accommodationEndpoint := fmt.Sprintf("%s:%s", server.config.AccommodationHost, server.config.AccommodationPort)
	fmt.Print("accommodationEndpoint: ")
	fmt.Println(accommodationEndpoint)
	err = accommodationGw.RegisterAccommodationServiceHandlerFromEndpoint(context.TODO(), server.mux, accommodationEndpoint, opts)
	if err != nil {
		panic(err)
	}

	authEndpoint := fmt.Sprintf("%s:%s", server.config.AuthHost, server.config.AuthPort)
	fmt.Print("authEndpoint: ")
	fmt.Println(authEndpoint)
	err = authGw.RegisterAuthServiceHandlerFromEndpoint(context.TODO(), server.mux, authEndpoint, opts)
	if err != nil {
		panic(err)
	}
	fmt.Println("Handlers initalized")

}

func (server *Server) initCustomHandlers() {
	catalogueEmdpoint := fmt.Sprintf("%s:%s", server.config.CatalogueHost, server.config.CataloguePort)
	orderingEmdpoint := fmt.Sprintf("%s:%s", server.config.OrderingHost, server.config.OrderingPort)
	shippingEmdpoint := fmt.Sprintf("%s:%s", server.config.ShippingHost, server.config.ShippingPort)
	orderingHandler := api.NewOrderingHandler(orderingEmdpoint, catalogueEmdpoint, shippingEmdpoint)
	orderingHandler.Init(server.mux)
}

func (server *Server) Start() {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), server.mux))
}
