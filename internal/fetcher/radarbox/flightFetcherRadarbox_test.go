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
			jsonFile:    "flights_test.json",
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
			jsonFile:    "empty_flights_test.json",
			expectedLen: 0,
			expectError: false,
		},
		{
			name:        "Invalid JSON format",
			jsonFile:    "invalid_flights_test.json",
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

			f := &radarbox.FlightsFetcherRadarbox{}
			flights := f.DecodeFlights(jsonResponse.Flights)
			if len(flights) != tt.expectedLen {
				t.Errorf("Expected %d flights, got %d", tt.expectedLen, len(flights))
			}

			if tt.expectedLen > 0 && flights[0] != tt.expectedFlight {
				t.Errorf("Expected flight %v, got %v", tt.expectedFlight, flights[0])
			}
		})
	}
}
