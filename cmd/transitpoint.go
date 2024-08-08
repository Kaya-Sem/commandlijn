package cmd

type GeoCoordinaat struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type TransitPoint struct {
	Entiteitnummer         string        `json:"entiteitnummer"`
	Haltenummer            string        `json:"haltenummer"`
	Omschrijving           string        `json:"omschrijving"`
	GeoCoordinaat          GeoCoordinaat `json:"geoCoordinaat"`
	HoofdHalte             *string       `json:"hoofdHalte"` // Use a pointer since it can be null
	HalteToegankelijkheden []string      `json:"halteToegankelijkheden"`
}
