package handlers_test

import (
	"flights/internal/fetcher"
	"flights/internal/handlers"
	"flights/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetFlights(t *testing.T) {
	mockFetcher := &fetcher.MockFlightsFetcher{
		FetchLatestFlightsFunc: func() ([]models.Flight, error) {
			return []models.Flight{
				{
					Airline:     "Mock Airline",
					Flight:      "MA123",
					Aircraft:    "Mock Aircraft",
					Origin:      "Mock Origin",
					Destination: "Mock Destination",
					HourDep:     "10:00",
					HourArr:     "12:00",
					Status:      "on time",
					TerminalDep: "A",
					TerminalArr: "B",
				},
			}, nil
		},
	}

	fhandler := handlers.FlightHandler{
		FlightsFetcher: mockFetcher,
	}

	req, err := http.NewRequest("GET", "/api/flights", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(fhandler.GetFlights)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `[{"airline":"Mock Airline","flight":"MA123","aircraft":"Mock Aircraft","origin":"Mock Origin","destination":"Mock Destination","hour_dep":"10:00","hour_arr":"12:00","status":"on time","terminal_dep":"A","terminal_arr":"B"}]`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetFlight(t *testing.T) {
	mockFetcher := &fetcher.MockFlightInfoFetcher{
		FetchFlightInfoFunc: func(flightID string) (models.Flight, error) {
			return models.Flight{
				Airline:     "Mock Airline",
				Flight:      flightID,
				Aircraft:    "Mock Aircraft",
				Origin:      "Mock Origin",
				Destination: "Mock Destination",
				HourDep:     "10:00",
				HourArr:     "12:00",
				Status:      "on time",
				TerminalDep: "A",
				TerminalArr: "B",
			}, nil
		},
	}

	fhandler := handlers.FlightHandler{
		FlightInfoFetcher: mockFetcher,
	}

	req, err := http.NewRequest("GET", "/api/flights/MA123", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/api/flights/{id}", fhandler.GetFlight).Methods("GET")
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"airline":"Mock Airline","flight":"MA123","aircraft":"Mock Aircraft","origin":"Mock Origin","destination":"Mock Destination","hour_dep":"10:00","hour_arr":"12:00","status":"on time","terminal_dep":"A","terminal_arr":"B"}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
