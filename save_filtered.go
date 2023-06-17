package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func (app *App) saveFilteredToCSV() error {
	//
	fileName := fmt.Sprintf("./filtered-%d.csv", time.Now().Unix())

	// Create a new CSV file in the current dir.
	file, err := os.Create(filepath.Clean(fileName))
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a new CSV writer instance.
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the header with PK column name of the CSV file.
	if err = writer.Write([]string{app.Config.ColumnsOrder[0]}); err != nil {
		return err
	}

	// Write data body of the CSV file.
	for _, data := range app.Outputs.Data {
		// Take only PK column data.
		if err = writer.Write([]string{data.ID}); err != nil {
			return err
		}
	}

	printStyled(
		fmt.Sprintf("– Saving filtered transactions to CSV file '%s' in the current dir... OK!", fileName),
		"margin-top",
	)

	return nil
}
