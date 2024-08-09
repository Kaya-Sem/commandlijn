package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// timetableCmd represents the timetable command
var timetableCmd = &cobra.Command{
	Use:   "timetable [transitpoint]",
	Short: "Fetch the timetable for a given transit point.",
	Long: `Fetch the timetable for a given transit point. By default, it uses the current time. 
You can specify a time using the -t or --time flag in 24-hour format (e.g., 1200 or 12:00). 
Additionally, you can use the optional flags:

  -d, --departure  : To get departure times (default behavior).
  -a, --arrival    : To get arrival times.

Examples:
  # Get the timetable for the current time and default to departure times.
  commandlijn timetable transitpoint123

  # Get the timetable for a specific time (e.g., 15:30) with departure times.
  commandlijn timetable transitpoint123 -t 15:30

  # Get the timetable for a specific time (e.g., 18:00) with arrival times.
  commandlijn timetable transitpoint123 -t 18:00 -a`,
	Args: cobra.ExactArgs(1), // Require exactly one argument
	Run: func(cmd *cobra.Command, args []string) {
		transitpoint := args[0] // Retrieve the transitpoint argument
		time, _ := cmd.Flags().GetString("time")
		isArrival, _ := cmd.Flags().GetBool("arrival")
		isDeparture, _ := cmd.Flags().GetBool("departure")

		// Default values
		if time == "" {
			time = "current time" // TODO: implement with actual time retrieval logic
		}

		if isArrival && isDeparture {
			fmt.Println("Please specify only one of --arrival or --departure.")
			return
		}

		if isArrival {
			fmt.Printf("Fetching arrival times for transit point %s at time %s\n", transitpoint, time)
		} else {
			fmt.Printf("Fetching departure times for transit point %s at time %s\n", transitpoint, time)
		}
	},
}

func init() {
	timetableCmd.Flags().StringP("time", "t", "", "Specify the time in 24-hour format (e.g., 1200 or 12:00)")
	timetableCmd.Flags().BoolP("arrival", "a", false, "Get arrival times")
	timetableCmd.Flags().BoolP("departure", "d", true, "Get departure times (default)")

	rootCmd.AddCommand(timetableCmd)
}
