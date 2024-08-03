package radarbox

import (
	"flights/models"
	"flights/utils"
)

func DecodeFlights(flightsJSON []map[string]interface{}) []models.Flight {
	var flights []models.Flight
	for _, flightJSON := range flightsJSON {
		flight := DecodeFlight(flightJSON)
		flights = append(flights, flight)
	}
	return flights
}

func DecodeFlight(flightJSON map[string]interface{}) models.Flight {
	return models.Flight{
		Airline:     utils.GetValueJson(flightJSON, "alna"),
		Flight:      utils.GetValueJson(flightJSON, "fnia"),
		Aircraft:    utils.GetValueJson(flightJSON, "act"),
		Origin:      utils.GetValueJson(flightJSON, "aporgci"),
		Destination: utils.GetValueJson(flightJSON, "apdstci"),
		HourDep:     utils.GetValueJson(flightJSON, "deps"),
		HourArr:     utils.GetValueJson(flightJSON, "arrs"),
		Status:      utils.GetValueJson(flightJSON, "status"),
	}
}
