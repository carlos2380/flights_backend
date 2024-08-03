package fetcher

import "flights/models"

type MockFlightsFetcher struct {
	FetchLatestFlightsFunc func() ([]models.Flight, error)
}

func (m *MockFlightsFetcher) FetchLatestFlights() ([]models.Flight, error) {
	return m.FetchLatestFlightsFunc()
}

type MockFlightInfoFetcher struct {
	FetchFlightInfoFunc func(flightID string) (models.Flight, error)
}

func (m *MockFlightInfoFetcher) FetchFlightInfo(flightID string) (models.Flight, error) {
	return m.FetchFlightInfoFunc(flightID)
}
