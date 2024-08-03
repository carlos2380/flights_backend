package radarbox

import (
	"encoding/json"
	"flights/internal/errors"
	"flights/models"
	"flights/utils"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FlightsFetcherRadarbox struct{}

const flightURL = "https://en.radarbox.com/search/active/flights?page=1"

func (f *FlightsFetcherRadarbox) FetchLatestFlights() ([]models.Flight, error) {
	resp, err := http.Get(flightURL)
	if err != nil {
		return nil, errors.DetailedError(errors.ErrFetchFlights, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.DetailedError(errors.ErrUnexpectedStatusCode, status.Errorf(codes.Internal, "unexpected status code: %s", http.StatusText(resp.StatusCode)))
	}

	var data struct {
		Flights []map[string]interface{} `json:"flights"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, errors.DetailedError(errors.ErrDecodeFlights, err)
	}

	flights := f.DecodeFlights(data.Flights)
	return flights, nil
}

func (f *FlightsFetcherRadarbox) DecodeFlights(flightsJSON []map[string]interface{}) []models.Flight {
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
