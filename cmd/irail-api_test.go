package cmd

import (
	"reflect"
	"testing"
)

// Test parseiRailTransitPoints with a valid JSON input
func TestParseiRailTransitPoints(t *testing.T) {
	jsonData := []byte(`{
		"version": "1.3",
		"timestamp": "1723234965",
		"station": [
			{
				"id": "BE.NMBS.008400319",
				"name": "'s Hertogenbosch",
				"standardname": "'s Hertogenbosch"
			},
			{
				"id": "BE.NMBS.008015345",
				"name": "Aachen Hbf",
				"standardname": "Aachen Hbf"
			}
		]
	}`)

	expectedTransitPoints := []TransitPoint{
		{
			Name:            "'s Hertogenbosch",
			Id:              "008400319",
			TransitProvider: string(SNCB),
			Description:     "",
		},
		{
			Name:            "Aachen Hbf",
			Id:              "008015345",
			TransitProvider: string(SNCB),
			Description:     "",
		},
	}

	transitPoints, err := parseiRailTransitPoints(jsonData)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(transitPoints, expectedTransitPoints) {
		t.Errorf("Expected %v, but got %v", expectedTransitPoints, transitPoints)
	}
}

// Test parseiRailTransitPoints with missing station ID digits

func TestParseiRailTransitPoints_InvalidIDFormat(t *testing.T) {
	jsonData := []byte(`{
		"version": "1.3",
		"timestamp": "1723234965",
		"station": [
			{
				"id": "BE.NMBS.XXXXXX",
				"name": "'s Hertogenbosch",
				"standardname": "'s Hertogenbosch"
			}
		]
	}`)

	_, err := parseiRailTransitPoints(jsonData)
	if err == nil {
		t.Fatal("Expected an error due to invalid ID format, but got nil")
	} else {
		t.Logf("Received expected error: %v", err)
	}
}