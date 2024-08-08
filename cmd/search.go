package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var limit int

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search [searchterm]",
	Short: "Search for public transport stops.",
	Long: `Search for public transport stops using a search term. 
The term can be a name of a place or stop.`,
	Args: cobra.ExactArgs(1), // Ensure exactly one argument is provided
	Run: func(cmd *cobra.Command, args []string) {
		searchterm := args[0] // Get the search term from the arguments
		haltesJson := apiHalteSearch(searchterm, limit)
		if haltesJson != nil {
			transitPoints, err := parseTransitPoints(haltesJson)
			if err != nil {
				fmt.Println("Error parsing JSON:", err)
				return
			}

			for _, tp := range transitPoints {
				fmt.Println(tp.Omschrijving, tp.Haltenummer)
			}
		}
	},
}

func parseTransitPoints(jsonData []byte) ([]TransitPoint, error) {
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

	// Return the slice of TransitPoint structs
	return result.Haltes, nil
}

func init() {
	// Define the flag for limit
	searchCmd.Flags().IntVarP(&limit, "limit", "l", 10, "Limit the number of results")
	rootCmd.AddCommand(searchCmd)
}
