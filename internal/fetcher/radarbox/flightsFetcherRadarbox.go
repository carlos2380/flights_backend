package radarbox

import (
	"encoding/json"
	"flights/internal/errors"
	"flights/models"
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

	flights := DecodeFlights(data.Flights)
	return flights, nil
}
