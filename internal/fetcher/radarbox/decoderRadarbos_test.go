package radarbox_test

import (
	"encoding/json"
	"flights/internal/fetcher/radarbox"
	"flights/models"
	"os"
	"testing"
)

func TestDecodeFlights(t *testing.T) {
	tests := []struct {
		name           string
		jsonFile       string
		expectedLen    int
		expectedFlight models.Flight
		expectError    bool
	}{
		{
			name:        "Valid JSON with 3 flights",
			jsonFile:    "data_test/flights_test.json",
			expectedLen: 3,
			expectError: false,
			expectedFlight: models.Flight{
				Airline:     "Qatar Airways",
				Flight:      "QR8909",
				Aircraft:    "B77L",
				Origin:      "Tokoname",
				Destination: "Mexico",
				HourDep:     "11:45",
				HourArr:     "",
				Status:      "departed",
			},
		},
		{
			name:        "Empty flights JSON",
			jsonFile:    "data_test/empty_flights_test.json",
			expectedLen: 0,
			expectError: false,
		},
		{
			name:        "Invalid JSON format",
			jsonFile:    "data_test/invalid_flights_test.json",
			expectedLen: 0,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonData, err := os.ReadFile(tt.jsonFile)
			if err != nil {
				t.Fatalf("Error reading JSON file: %v", err)
			}

			var jsonResponse struct {
				Flights []map[string]interface{} `json:"flights"`
			}
			err = json.Unmarshal(jsonData, &jsonResponse)
			if (err != nil) != tt.expectError {
				t.Fatalf("Expected error: %v, got: %v", tt.expectError, err)
			}

			if err != nil {
				return
			}

			flights := radarbox.DecodeFlights(jsonResponse.Flights)
			if len(flights) != tt.expectedLen {
				t.Errorf("Expected %d flights, got %d", tt.expectedLen, len(flights))
			}

			if tt.expectedLen > 0 && flights[0] != tt.expectedFlight {
				t.Errorf("Expected flight %v, got %v", tt.expectedFlight, flights[0])
			}
		})
	}
}

func TestDecodeFlight(t *testing.T) {
	tests := []struct {
		name           string
		jsonFile       string
		expectedFlight models.Flight
		expectError    bool
	}{
		{
			name:        "Valid JSON for single flight",
			jsonFile:    "data_test/single_flight_test.json",
			expectError: false,
			expectedFlight: models.Flight{
				Airline:     "American Airlines",
				Flight:      "AA980",
				Aircraft:    "A321",
				Origin:      "Atlanta, GA",
				Destination: "Dallas-Fort Worth, TX",
				HourDep:     "14:36",
				HourArr:     "15:53",
				Status:      "live",
			},
		},
		{
			name:        "Invalid JSON format",
			jsonFile:    "data_test/invalid_single_flight_test.json",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonData, err := os.ReadFile(tt.jsonFile)
			if err != nil {
				t.Fatalf("Error reading JSON file: %v", err)
			}

			var flightJSON map[string]interface{}
			err = json.Unmarshal(jsonData, &flightJSON)
			if (err != nil) != tt.expectError {
				t.Fatalf("Expected error: %v, got: %v", tt.expectError, err)
			}

			if err != nil {
				return
			}

			flight := radarbox.DecodeFlight(flightJSON)
			if flight != tt.expectedFlight {
				t.Errorf("Expected flight %v, got %v", tt.expectedFlight, flight)
			}
		})
	}
}
