package cmd

import (
	"fmt"
	"strings"
)

func replaceSpacesWithURLCode(input string) string {
	return strings.ReplaceAll(input, " ", "%20")
}

func printTransitPoints(tp []TransitPoint) {
	for _, t := range tp {
		fmt.Printf("Entiteitnummer: %s\nOmschrijving: %s\nHaltenummer: %s\nProvider: %s\n\n",
			t.Entiteitnummer, t.Omschrijving, t.Haltenummer, t.TransitProvider)
	}
}
