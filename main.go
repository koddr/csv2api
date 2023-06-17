package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Listen to flags.
	flag.BoolVar(&initConfig, "i", false, "set to generate initial config and example data files")
	flag.StringVar(&configFilePath, "c", "./config.yml", "set path to your config file")
	flag.StringVar(&inputDataFilePath, "d", "./data.csv", "set path to your CSV file with input data")
	flag.StringVar(&envPrefix, "e", "CONFIG", "set prefix used in your environment variables")

	// Parse all given flags.
	flag.Parse()

	// If '-i' flag is true, generates config and example data files.
	if initConfig {
		// Create the config file in the current dir.
		if err := os.WriteFile(filepath.Clean("./config.yml"), embedConfigYAMLFile, 0o644); err != nil {
			log.Fatal(err)
		}

		// Create the example data file in the current dir.
		if err := os.WriteFile(filepath.Clean("./data.csv"), embedDataCSVFile, 0o644); err != nil {
			log.Fatal(err)
		}

		// Show success message.
		log.Println("The configuration and example data files was successfully generated in the current dir!")
	} else {
		// App initialization.
		app, err := initialize()
		if err != nil {
			log.Fatal(err)
		}

		// App start.
		if err = app.start(); err != nil {
			log.Fatal(err)
		}
	}
}
