package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func addDeLijnHeaderMetadata(request *http.Request) {
	apiKey := GetConfig().DeLijnAPIKey // TODO: globally define config
	request.Header.Add("Cache-Control", "no-cache")
	request.Header.Add("Ocp-Apim-Subscription-Key", apiKey)
}

// getDeLijnHaltesJSON performs the API request and returns the response body as JSON (byte slice).
func getDeLijnHaltesJSON(searchterm string, limit int) []byte {
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

	// TODO:
	//logVerbose(fmt.Sprintf("\nStatus code: %s", StatusCodes[resp.StatusCode]))
	return body
}

type Halte struct {
	Entiteitnummer string `json:"entiteitnummer"`
	Haltenummer    string `json:"haltenummer"`
	Omschrijving   string `json:"omschrijving"`
}

type ApiResponse struct {
	AantalHits int     `json:"aantalHits"`
	Haltes     []Halte `json:"haltes"`
}

func parseDeLijnTransitPoints(jsonData []byte) ([]Halte, error) {
	var result ApiResponse

	err := json.Unmarshal(jsonData, &result)
	if err != nil {
		return nil, err
	}

	// TODO: find a solution to not disrupt spinner
	// logVerbose(fmt.Sprintf("total results: %d\n", result.AantalHits))
	return result.Haltes, nil
}

// https://portal.delijn.be/api-details#api=KernOpenDataServicesV1&operation=get-haltes-entiteitnummer-haltenummer-dienstregelingen
func getDeLijnHalteTimeTable(entityID string, halteID string) ([]byte, error) {
	url := fmt.Sprintf("https://api.delijn.be/DLKernOpenData/api/v1/haltes/%s/%s/dienstregelingen", entityID, halteID)

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
