package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Listen to flags.
	flag.BoolVar(&initConfig, "i", false, "set to generate initial config and example data files")
	flag.StringVar(&configFilePath, "c", "", "set path to your config file")
	flag.StringVar(&inputDataFilePath, "d", "", "set path to your CSV file with input data")
	flag.StringVar(&envPrefix, "e", "CONFIG", "set prefix used in your environment variables")

	// Parse all given flags.
	flag.Parse()

	// If '-i' flag is true, generates config and example data files.
	if initConfig {
		// Create the config file in the current dir.
		if err := os.WriteFile(filepath.Clean("./config.yml"), embedConfigYAMLFile, 0o600); err != nil {
			printStyled(
				fmt.Sprintf("✕ There was an error with generate config.yml file: %v", err),
				"",
			)
		}

		// Create the example data file in the current dir.
		if err := os.WriteFile(filepath.Clean("./data.csv"), embedDataCSVFile, 0o600); err != nil {
			printStyled(
				fmt.Sprintf("✕ There was an error with generate data.csv file: %v", err),
				"",
			)
		}

		// Show success message.
		printStyled(
			"The configuration and example data files was successfully generated in the current dir!",
			"margin-top-bottom",
		)
	} else {
		// App initialization.
		app, err := initialize()
		if err != nil {
			printStyled(
				fmt.Sprintf("✕ There was an error with initialize app: %v", err),
				"",
			)
		}

		// App start.
		if err = app.start(); err != nil {
			printStyled(
				fmt.Sprintf("✕ There was an error with starting app: %v", err),
				"",
			)
		}
	}
}
