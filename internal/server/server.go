package server

import (
	"flights/internal/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(fHandler *handlers.FlightHandler) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/api/flights", fHandler.GetFlights).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/flights/{id}", fHandler.GetFlight).Methods(http.MethodGet, http.MethodOptions)
	return r
}
