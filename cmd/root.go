package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Version string // This will be set from main.go

var (
	verbose bool
	rootCmd = &cobra.Command{
		Use:   "commandlijn",
		Short: "A brief description of your application",
		Long:  ``,
		// Define the default action when no subcommands are provided
		Run: func(cmd *cobra.Command, args []string) {
			if v, _ := cmd.Flags().GetBool("version"); v {
				fmt.Println("Version:", Version)
				os.Exit(0)
			}
			fmt.Println("Use --help to see available commands.")
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Define the --version flag
	rootCmd.Flags().BoolP("version", "V", false, "Print the version number and exit")

	// Define other flags
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
