package main

import (
	"fmt"
	"time"
)

// newApp provides a new application instance.
func newApp(config *Config, inputs *Inputs, filtered *Filtered, outputs *Outputs) *App {
	return &App{
		Config:   config,
		Inputs:   inputs,
		Filtered: filtered,
		Outputs:  outputs,
	}
}

// start starts an application with a beauty output to console.
func (app *App) start() error {
	// Start timer.
	start := time.Now()

	// Check, if input data is not empty.
	if len(app.Inputs.Data) > 0 {
		printStyled(
			"Hello and welcome to csv2api! 👋",
			"margin-top-bottom",
		)
		printStyled(
			fmt.Sprintf(
				"– According to the settings in '%s', %d transactions were filtered out of %d to start the process... Please wait!",
				configFilePath, len(app.Filtered.Data), len(app.Inputs.Data),
			),
			"margin-bottom",
		)

		// Check, if output data is not empty.
		if len(app.Outputs.Data) > 0 {
			// Loop for all output data.
			for _, data := range app.Outputs.Data {
				// Start updating fields.
				if err := app.updateField(data.ID, data.FieldName, data.Values); err != nil {
					printStyled(
						fmt.Sprintf("✕ There was an error: %v", err),
						"margin-left",
					)
				}
			}

			// Check, if you need to save CSV file with filtered PK.
			if app.Config.SaveFilteredPKToCSV {
				// Save filtered PK to CSV.
				if err := app.saveFilteredPKToCSV(); err != nil {
					printStyled(
						fmt.Sprintf("✕ There was an error: %v", err),
						"margin-top",
					)
				}
			}
		}
	}

	printStyled(
		fmt.Sprintf(
			"All done! 🎉 Time elapsed: %.2fs",
			time.Since(start).Seconds(),
		),
		"margin-top-bottom",
	)

	return nil
}
