package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func (fhandler FlightHandler) GetFlights(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	switch r.Method {
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
	case http.MethodGet:
		retStores, err := fhandler.FlightFetcher.FetchLatestFlights()
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(
				struct {
					Error string `json:"error"`
				}{Error: "Bad Request"})
			return
		}
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(retStores)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
