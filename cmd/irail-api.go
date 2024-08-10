package cmd

// https://docs.irail.be/

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	BaseURL             = "https://api.irail.be"
	allStationsURL      = BaseURL + "/stations/?format=json&lang=en"
	StationTimetableURL = BaseURL + "/liveboard/?id=%s&station=%s&time=%s&arrdep=%s&lang=nl&format=json"
)

// https://docs.irail.be/#liveboard-liveboard-api-get
func getSNCBStationTimeTable(stationId string, time string, arrdep string) ([]byte, error) {
	url := fmt.Sprintf(StationTimetableURL, stationId, "", time, arrdep)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Parse the iRail departures JSON into a slice of Departure structs
func parseiRailDepartures(jsonData []byte) ([]Departure, error) {
	var response StationTimetableResponse
	err := json.Unmarshal(jsonData, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v - input data: %s", err, string(jsonData))
	}

	// Extract the list of departures
	return response.Departures.Departure, nil
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

	// Convert the parsed data to the TransitPoint struct
	transitPoints := make([]TransitPoint, len(result.Stations))
	for i, station := range result.Stations {
		transitPoints[i] = TransitPoint{
			Name:            station.Name,
			Id:              station.ID,
			TransitProvider: string(SNCB),
			Description:     "",
		}
	}

	return transitPoints, nil
}
