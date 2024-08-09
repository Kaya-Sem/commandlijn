package cmd

import (
	"fmt"
	"strings"
	"time"
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

func getCurrentTimeHHMM() string {
	hours, minutes, _ := time.Now().Clock()
	return fmt.Sprintf("%d%02d", hours, minutes)

}
func normalizeTime(time string) string {
	time = strings.ReplaceAll(time, " ", "")
	time = strings.ReplaceAll(time, ":", "")
	return time
}
