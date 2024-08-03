package fetcher

import "flights/models"

type FlightInfoFetcher interface {
	FetchFlightInfo(flight string) (models.Flight, error)
}
