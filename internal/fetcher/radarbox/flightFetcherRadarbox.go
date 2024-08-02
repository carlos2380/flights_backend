package radarbox

import (
	"encoding/json"
	"flights/models"
	"flights/utils"
	"net/http"
)

type FlightFetcherRadarbox struct{}

const flightURL = "https://en.radarbox.com/search/active/flights?page=1"

func (f *FlightFetcherRadarbox) FetchLatestFlights() ([]models.Flight, error) {
	resp, err := http.Get(flightURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		Flights []map[string]interface{} `json:"flights"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	flights := f.DecodeFlights(data.Flights)
	return flights, nil
}

func (f *FlightFetcherRadarbox) DecodeFlights(flightsJSON []map[string]interface{}) []models.Flight {
	var flights []models.Flight
	for _, flightJSON := range flightsJSON {
		flight := models.Flight{
			Airline:     utils.GetValueJson(flightJSON, "alna"),
			Flight:      utils.GetValueJson(flightJSON, "fnia"),
			Aircraft:    utils.GetValueJson(flightJSON, "act"),
			Origin:      utils.GetValueJson(flightJSON, "aporgci"),
			Destination: utils.GetValueJson(flightJSON, "apdstci"),
			HourDep:     utils.GetValueJson(flightJSON, "deps"),
			HourArr:     utils.GetValueJson(flightJSON, "arrs"),
			Status:      utils.GetValueJson(flightJSON, "status"),
		}
		flights = append(flights, flight)
	}
	return flights
}
