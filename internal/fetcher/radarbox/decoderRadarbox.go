package radarbox

import (
	"flights/models"
	"flights/utils"
	"fmt"
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
		HourDep:     fmt.Sprintf("%s %s", utils.GetValueJson(flightJSON, "deps"), utils.GetValueJson(flightJSON, "aporgtzns")),
		HourArr:     fmt.Sprintf("%s %s", utils.GetValueJson(flightJSON, "arrs"), utils.GetValueJson(flightJSON, "apdsttzns")),
		Status:      fmt.Sprintf("%s %s %s", utils.GetValueJson(flightJSON, "status"), utils.GetValueJson(flightJSON, "arre"), utils.GetValueJson(flightJSON, "apdsttzns")),
		TerminalDep: utils.GetValueJson(flightJSON, "depterm"),
		TerminalArr: utils.GetValueJson(flightJSON, "arrterm"),
	}
}
