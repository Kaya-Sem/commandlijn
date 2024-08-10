package cmd

type TransitProvider string

const (
	DELIJN TransitProvider = "De Lijn"
	SNCB   TransitProvider = "SNCB"
)

// TransitPoint is a general representation of a stop or station
type TransitPoint struct {
	Name            string
	Id              string
	TransitProvider string
	Description     string
}
