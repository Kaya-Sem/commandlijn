package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	verbose bool
	rootCmd = &cobra.Command{
		Use:   "commandlijn",
		Short: "A brief description of your application",
		Long:  ``,
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
