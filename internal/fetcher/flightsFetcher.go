package fetcher

import "flights/models"

type FlightsFetcher interface {
	FetchLatestFlights() ([]models.Flight, error)
}
