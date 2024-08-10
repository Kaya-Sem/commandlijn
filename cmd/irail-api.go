package cmd

// https://docs.irail.be/

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

const (
	allStationsURL = "https://api.irail.be/stations/?format=json&lang=en"
)

// TODO:
func getSNCBStationTimeTable(name string) []byte {
	return nil
}

func getSNCBStationsJSON() []byte {
	req, err := http.NewRequest("GET", allStationsURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil
	}

	logVerbose(fmt.Sprintf("\nStatus code: %s", StatusCodes[resp.StatusCode]))
	return body
}

// Intermediate struct to match the JSON structure
type stationJSON struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	StandardName string `json:"standardname"`
}

var result struct {
	Stations []stationJSON `json:"station"`
}

// https://docs.irail.be/#stations-stations-api-get

func parseiRailTransitPoints(jsonData []byte) ([]TransitPoint, error) {
	err := json.Unmarshal(jsonData, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v - input data: %s", err, string(jsonData))
	}

	idRegex := regexp.MustCompile(`\d+`)

	// Convert the parsed data to the TransitPoint struct
	transitPoints := make([]TransitPoint, len(result.Stations))
	for index, station := range result.Stations {
		numericID := idRegex.FindString(station.ID)

		// Check if numericID is empty, meaning no digits were found
		if numericID == "" {
			return nil, fmt.Errorf("invalid ID format for station %s: no numeric part found in ID '%s'", station.Name, station.ID)
		}

		transitPoints[index] = TransitPoint{
			Name:            station.Name,
			Id:              numericID,
			TransitProvider: string(SNCB),
			Description:     "",
		}
	}

	return transitPoints, nil
}
