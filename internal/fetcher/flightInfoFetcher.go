package fetcher

import "flights/models"

type FlightInfoFetcher interface {
	FetchFlightInfo() ([]models.Flight, error)
}
