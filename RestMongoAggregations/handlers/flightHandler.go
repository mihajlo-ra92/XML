package handlers

import (
	"Rest/data"
	"context"
	"log"
	"net/http"

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

func (u *FlightsHandler) InitTestDb(rw http.ResponseWriter, h *http.Request) {
	err := u.repo.DropCollection()
	if err != nil {
		u.logger.Print("Database exception: ", err)
	}
}

func (f *FlightsHandler) PostFlight(rw http.ResponseWriter, h *http.Request) {
	flight := h.Context().Value(KeyProduct{}).(*data.Flight)
	f.repo.Insert(flight)
	rw.WriteHeader(http.StatusCreated)
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
