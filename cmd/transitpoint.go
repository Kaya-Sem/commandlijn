package cmd

type TransitProvider string

const (
	DELIJN TransitProvider = "De Lijn"
	SNCB   TransitProvider = "SNCB"
)

type TransitPoint struct {
	Name            string
	Id              string
	TransitProvider string
	Description     string
}
