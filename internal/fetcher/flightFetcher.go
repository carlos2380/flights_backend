package fetcher

import "flights/models"

type FlightFetcher interface {
	FetchLatestFlights() ([]models.Flight, error)
}
