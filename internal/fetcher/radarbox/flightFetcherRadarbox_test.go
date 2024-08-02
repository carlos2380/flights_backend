package radarbox_test

import (
	"encoding/json"
	"flights/internal/fetcher/radarbox"
	"flights/models"
	"os"
	"testing"
)

func TestDecodeFlights(t *testing.T) {
	// Leer el contenido del archivo JSON
	jsonData, err := os.ReadFile("flights_test.json")
	if err != nil {
		t.Fatalf("Error reading JSON file: %v", err)
	}

	// Deserializar el JSON en la estructura adecuada
	var jsonResponse struct {
		Flights []map[string]interface{} `json:"flights"`
	}
	if err := json.Unmarshal(jsonData, &jsonResponse); err != nil {
		t.Fatalf("Error unmarshalling JSON: %v", err)
	}

	// Verificar si hay vuelos en la respuesta JSON
	if len(jsonResponse.Flights) == 0 {
		t.Fatal("No flights found in JSON response")
	}

	// Usar el método DecodeFlights del paquete radarbox
	f := &radarbox.FlightFetcherRadarbox{}
	flights := f.DecodeFlights(jsonResponse.Flights)
	if len(flights) != 3 {
		t.Errorf("Expected 3 flights, got %d", len(flights))
	}

	// Definir el vuelo esperado
	expected := models.Flight{
		Airline:     "Qatar Airways",
		Flight:      "QR8909",
		Aircraft:    "B77L",
		Origin:      "Tokoname",
		Destination: "Mexico",
		HourDep:     "11:45",
		HourArr:     "",
		Status:      "departed",
	}

	// Comparar el primer vuelo decodificado con el vuelo esperado
	if flights[0] != expected {
		t.Errorf("Expected %v, got %v", expected, flights[0])
	}
}
