package radarbox

import (
	"encoding/json"
	"flights/internal/errors"
	"flights/models"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FlightInfoRadarbox struct{}

const flightInfoURL = "https://en.radarbox.com/data/flights/info?type=flights&query="

func (f *FlightInfoRadarbox) FetchFlightInfo(flight string) (models.Flight, error) {
	url := flightInfoURL + flight
	resp, err := http.Get(url)
	if err != nil {
		return models.Flight{}, errors.DetailedError(errors.ErrFetchFlightInfo, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.Flight{}, errors.DetailedError(errors.ErrUnexpectedStatusCode, status.Errorf(codes.Internal, "unexpected status code: %s", http.StatusText(resp.StatusCode)))
	}

	if resp.ContentLength == 0 {
		return models.Flight{}, nil
	}

	var flightJSON map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&flightJSON); err != nil {
		return models.Flight{}, errors.DetailedError(errors.ErrDecodeFlightInfo, err)
	}

	flightInfo := DecodeFlight(flightJSON)
	return flightInfo, nil
}
