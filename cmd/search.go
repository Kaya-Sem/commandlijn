package cmd

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
)

// Define the default search limit
const SearchLimit int = 10

var limit int

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search [searchterm]",
	Short: "Search for public transport stops.",
	Long: `Search for public transport stops using a search term. 
The term can be a name of a place or stop. Searches for both De Lijn and SNCB`,
	Args: cobra.ExactArgs(1), // Ensure exactly one argument is provided
	Run: func(cmd *cobra.Command, args []string) {
		searchterm := args[0] // Get the search term from the arguments
		trainSet := []string{"🚅", "🚆", "🚈", "🚌"}
		s := spinner.New(trainSet, 100*time.Millisecond)
		//s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Start()
		defer s.Stop()

		// TODO: also search SNCB stops
		haltesJson := apiHalteSearch(searchterm, limit)
		s.Stop()
		if haltesJson != nil {
			transitPoints, err := parseDeLijnTransitPoints(haltesJson)
			if err != nil {
				fmt.Println("Error parsing JSON:", err)
				return
			}

			printTransitPoints(transitPoints)
		}
	},
}

func init() {
	// Define the flag for limit
	searchCmd.Flags().IntVarP(&limit, "limit", "l", SearchLimit, "Limit the number of results")
	rootCmd.AddCommand(searchCmd)
}
