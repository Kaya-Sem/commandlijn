package cmd

import "encoding/json"

type TransitProvider string

const (
	DELIJN TransitProvider = "De Lijn"
	SNCB   TransitProvider = "SNCB"
)

type GeoCoordinaat struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type TransitPoint struct {
	TransitProvider        TransitProvider // not filled by JSON
	GeoCoordinaat          GeoCoordinaat   `json:"geoCoordinaat"`
	HalteToegankelijkheden []string        `json:"halteToegankelijkheden"`
	HoofdHalte             *string         `json:"hoofdHalte"` // Use a pointer since it can be null
	Entiteitnummer         string          `json:"entiteitnummer"`
	Haltenummer            string          `json:"haltenummer"`
	Omschrijving           string          `json:"omschrijving"`
}

func NewTransitPoint(
	entiteitnummer string,
	haltenummer string,
	omschrijving string,
	geoCoordinaat GeoCoordinaat,
	hoofdHalte *string,
	halteToegankelijkheden []string,
	transitProvider TransitProvider,
) *TransitPoint {
	return &TransitPoint{
		Entiteitnummer:         entiteitnummer,
		Haltenummer:            haltenummer,
		Omschrijving:           omschrijving,
		GeoCoordinaat:          geoCoordinaat,
		HoofdHalte:             hoofdHalte,
		HalteToegankelijkheden: halteToegankelijkheden,
		TransitProvider:        transitProvider,
	}
}

func UnmarshalTransitPoint(data []byte) (*TransitPoint, error) {
	var tp TransitPoint
	if err := json.Unmarshal(data, &tp); err != nil {
		return nil, err
	}
	return &tp, nil
}
