package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func addDeLijnHeaderMetadata(request *http.Request) {
	apiKey := os.Getenv("API_KEY") // TODO: handle this in config file
	request.Header.Add("Cache-Control", "no-cache")
	request.Header.Add("Ocp-Apim-Subscription-Key", apiKey)
}

// apiHalteSearch performs the API request and returns the response body as JSON (byte slice).
func apiHalteSearch(searchterm string, limit int) []byte {
	searchterm = replaceSpacesWithURLCode(searchterm)

	url := fmt.Sprintf("https://api.delijn.be/DLZoekOpenData/v1/zoek/haltes/%s?startIndex=0&maxAantalHits=%d", searchterm, limit)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil
	}

	addDeLijnHeaderMetadata(req)

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

	logVerbose(fmt.Sprintf("\nStatus code: %s\n", StatusCodes[resp.StatusCode]))
	return body
}

func parseDeLijnTransitPoints(jsonData []byte) ([]TransitPoint, error) {
	// Create a struct to match the top-level JSON structure
	var result struct {
		AantalHits int            `json:"aantalHits"`
		Haltes     []TransitPoint `json:"haltes"`
	}

	// Unmarshal the JSON into the struct
	err := json.Unmarshal(jsonData, &result)
	if err != nil {
		return nil, err
	}

	// Manually set the TransitProvider for each TransitPoint
	for i := range result.Haltes {
		result.Haltes[i].TransitProvider = DELIJN
	}

	return result.Haltes, nil
}
