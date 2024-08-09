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
		fmt.Printf("%s Haltenummer: %s Provider: %s\n", t.Omschrijving, t.Haltenummer, t.TransitProvider)
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

func logVerbose(msg string) {
	if verbose {
		fmt.Println(msg)
	}
}
