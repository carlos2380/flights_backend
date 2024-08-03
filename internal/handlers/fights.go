package handlers

import (
	"encoding/json"
	"flights/internal/errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// GetFlights godoc
// @Summary Get last flights
// @Description Get the list of latest flights
// @Tags flights
// @Produce  json
// @Success 200 {array} models.Flight
// @Failure 500 {object} map[string]string
// @Router /api/flights [get]
func (fhandler FlightHandler) GetFlights(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	switch r.Method {
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
	case http.MethodGet:
		retStores, err := fhandler.FlightsFetcher.FetchLatestFlights()
		if err != nil {
			log.Println(err)
			errors.WriteJSONError(w, err, http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(retStores)
	default:
		errors.WriteJSONError(w, errors.ErrMethodNotAllowed, http.StatusMethodNotAllowed)
	}
}

// GetFlight godoc
// @Summary Get flight by id flight
// @Description Get details of a flight by its id flight
// @Tags flights
// @Produce  json
// @Param id path string true "Flight ID"
// @Success 200 {object} models.Flight
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/flights/{id} [get]
func (fhandler FlightHandler) GetFlight(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	switch r.Method {
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
	case http.MethodGet:
		vars := mux.Vars(r)
		flightID := vars["id"]
		if flightID == "" {
			errors.WriteJSONError(w, errors.ErrInvalidFlightID, http.StatusBadRequest)
			return
		}
		flightInfo, err := fhandler.FlightInfoFetcher.FetchFlightInfo(flightID)
		if err != nil {
			log.Println(err)
			errors.WriteJSONError(w, err, http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(flightInfo)
	default:
		errors.WriteJSONError(w, errors.ErrMethodNotAllowed, http.StatusMethodNotAllowed)
	}
}
