package handlers

import (
	"encoding/json"
	"flights/internal/errors"
	"log"
	"net/http"
)

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
