package cmd

import (
	"fmt"
	"strings"
	"time"
)

const (
	StatusOK                  = 200
	StatusInternalServerError = 500
)

var StatusCodes = map[int]string{
	StatusOK:                  "\033[32m200 OK\033[0m",                    // green
	StatusInternalServerError: "\033[31m500 Internal Server Error\033[0m", // red
}

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
