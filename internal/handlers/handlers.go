package handlers

import (
	"flights/internal/fetcher"
	"net/http"
)

type FlightHandler struct {
	FlightsFetcher    fetcher.FlightsFetcher
	FlightInfoFetcher fetcher.FlightInfoFetcher
}

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
}
