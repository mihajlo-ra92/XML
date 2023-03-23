package handlers

import (
	"Rest/data"
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type TicketsHandler struct {
	logger     *log.Logger
	repo       *data.TicketRepo
	flightRepo *data.FlightRepo
}

func NewTicketsHandler(l *log.Logger, r *data.TicketRepo, f *data.FlightRepo) *TicketsHandler {
	return &TicketsHandler{l, r, f}
}

func (u *TicketsHandler) MiddlewareTicketDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		ticket := &data.Ticket{}
		err := ticket.FromJSON(h.Body)
		u.logger.Println(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			u.logger.Fatal(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProduct{}, ticket)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}

func (u *TicketsHandler) InitTestDb(rw http.ResponseWriter, h *http.Request) {
	err := u.repo.DropCollection()
	if err != nil {
		u.logger.Print("Database exception: ", err)
	}
}

func (t *TicketsHandler) PostTicket(rw http.ResponseWriter, h *http.Request) {
	ticket := h.Context().Value(KeyProduct{}).(*data.Ticket)
	flight, err := t.flightRepo.GetFlightById(ticket.FlightId)
	if err != nil {
		t.logger.Print("Exception: ", err)
		return
	}
	ticketState := t.repo.CheckSeats(ticket, flight[0])
	if ticketState == false {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	flight[0].FreeSeats = flight[0].FreeSeats - ticket.Capacity
	t.flightRepo.Update(ticket.FlightId, flight[0])
	t.repo.Insert(ticket)
	rw.WriteHeader(http.StatusCreated)
}

func (f *TicketsHandler) GetAllTickets(rw http.ResponseWriter, h *http.Request) {
	tickets, err := f.repo.GetAll()
	if err != nil {
		f.logger.Print("Database exception: ", err)
	}

	if tickets == nil {
		return
	}

	err = tickets.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		f.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (t *TicketsHandler) GetTicketById(rw http.ResponseWriter, h *http.Request) {
	id := h.URL.Query().Get("id")

	ticket, err := t.repo.GetById(id)
	if err != nil {
		t.logger.Print("Database exception: ", err)
	}

	if ticket == nil {
		return
	}

	err = ticket.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		t.logger.Fatal("Unable to convert to json: ", err)
		return
	}
}

func (t *TicketsHandler) GetTicketsByUserId(rw http.ResponseWriter, h *http.Request) {
	id := h.URL.Query().Get("id")

	tickets, err := t.repo.GetByIUserId(id)
	if err != nil {
		t.logger.Print("Database exception: ", err)
	}

	if tickets == nil {
		return
	}

	err = tickets.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		t.logger.Fatal("Unable to convert to json: ", err)
		return
	}
}

func (t *TicketsHandler) PatchTicket(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]
	ticket := h.Context().Value(KeyProduct{}).(*data.Ticket)

	t.repo.Update(id, ticket)
	rw.WriteHeader(http.StatusOK)
}

func (t *TicketsHandler) DeleteTicket(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	id := vars["id"]

	t.repo.Delete(id)
	rw.WriteHeader(http.StatusNoContent)
}
