package main

import (
	"fmt"
	"time"
)

// newApp ...
func newApp(config *Config, inputs *Inputs, filtered *Filtered, outputs *Outputs) *App {
	return &App{
		Config:   config,
		Inputs:   inputs,
		Filtered: filtered,
		Outputs:  outputs,
	}
}

// start ...
func (app *App) start() error {
	//
	start := time.Now()

	//
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

		//
		if len(app.Outputs.Data) > 0 {
			//
			for _, data := range app.Outputs.Data {
				//
				if err := app.updateField(data.ID, data.FieldName, data.Values); err != nil {
					printStyled(
						fmt.Sprintf("✕ There was an error: %v", err),
						"margin-left",
					)
				}
			}

			//
			if app.Config.SaveFilteredToCSV {
				//
				if err := app.saveFilteredToCSV(); err != nil {
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
