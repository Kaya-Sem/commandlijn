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
	Use:   "search",
	Short: "Search for public transport stops.",
	Long: `Search for public transport stops using a search term. 
The term can be a name of a place or stop. Searches for both De Lijn and SNCB.`,
}

var delijnCmd = &cobra.Command{
	Use:   "delijn [searchterm]",
	Short: "Search for De Lijn public transport stops.",
	Long:  `Search for De Lijn public transport stops using a search term. The term can be a name of a place or stop.`,
	Args:  cobra.ExactArgs(1), // Ensure exactly one argument is provided
	Run: func(cmd *cobra.Command, args []string) {
		searchterm := args[0] // Get the search term

		s := spinner.New(spinner.CharSets[35], 250*time.Millisecond)
		s.Prefix = "searching stops "
		s.Start()
		defer s.Stop()
		time.Sleep(1 * time.Second)

		haltesJson := getDeLijnHaltesJSON(searchterm, limit)

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

// sncbCmd represents the sncb/nmbs subcommand
var sncbCmd = &cobra.Command{
	Use:   "sncb [searchterm]",
	Short: "Search for SNCB/NMBS public transport stops.",
	Long:  `Search for SNCB/NMBS public transport stops using a search term. The term can be a name of a place or stop.`,
	Args:  cobra.ExactArgs(1), // Ensure exactly one argument is provided
	Run: func(cmd *cobra.Command, args []string) {
		searchterm := args[0] // Get the search term from the arguments

		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Suffix = " searching stations..."
		s.Start()
		defer s.Stop()
		time.Sleep(1 * time.Second)

		// TODO: Implement SNCB/NMBS search logic
		fmt.Printf("todo implement searchterm filtering: %s", searchterm)
		sncbJson := getSNCBStationsJSON()
		s.Stop()
		if sncbJson != nil {
			transitPoints, err := parseiRailTransitPoints(sncbJson)
			if err != nil {
				fmt.Println("Error parsing JSON:", err)
				return
			}

			printTransitPoints(transitPoints)
		}
	},
}

func init() {
	delijnCmd.Flags().IntVarP(&limit, "limit", "l", SearchLimit, "Limit the number of results")

	// Add subcommands to the search command
	searchCmd.AddCommand(delijnCmd)
	searchCmd.AddCommand(sncbCmd)

	rootCmd.AddCommand(searchCmd)
}
