package cmd

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the commandlijn configuration file",
	Long: `This command initializes a configuration file at ~/.config/commandlijn/commandlijn.yaml
with the required API key for DeLijn.`,
	Run: func(cmd *cobra.Command, args []string) {

		// 1) check if config already exists
		if _, err := os.Stat(getConfigFilePath()); err == nil {
			fmt.Printf("Config is already present at %s", getConfigFilePath())
			os.Exit(0)
		}

		if err := initializeConfig(); err != nil {
			fmt.Println("Error initializing configuration:", err)
		} else {
			fmt.Println("\nConfiguration file initialized successfully.")
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func getConfigDir() string {
	return filepath.Join(os.Getenv("HOME"), ".config", "commandlijn")
}

func getConfigFilePath() string {
	configDir := getConfigDir()
	configFile := filepath.Join(configDir, "commandlijn.yaml")
	return configFile
}

// initializeConfig handles the creation of the configuration file
func initializeConfig() error {
	configFile := getConfigFilePath()

	// Ensure the config directory exists
	if err := os.MkdirAll(getConfigDir(), 0700); err != nil {
		return fmt.Errorf("error creating config directory: %v", err)
	}

	// Prompt user for API keys
	delijnAPIKey, err := promptForInput("Enter DeLijn API key:")
	if err != nil {
		return fmt.Errorf("error reading DeLijn API key: %v", err)
	}

	// Create configuration
	config := Config{
		DeLijnAPIKey: delijnAPIKey,
		Aliases: []Alias{
			{
				Name:     "GSP",
				Provider: SNCB,
				ID:       []string{"BE.NMBS.008892007"},
			},
		},
	}

	data, err := yaml.Marshal(&config)
	if err != nil {
		return fmt.Errorf("error marshaling YAML: %v", err)
	}

	if err := os.WriteFile(configFile, data, 0600); err != nil {
		return fmt.Errorf("error writing config file: %v", err)
	}

	return nil
}

// promptForInput prompts the user for input and returns the result
func promptForInput(label string) (string, error) {
	prompt := promptui.Prompt{
		Label: label,
	}

	result, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("prompt failed: %v", err)
	}

	return result, nil
}
