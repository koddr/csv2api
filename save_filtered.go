package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// saveFilteredPKToCSV provides save process to filtered PK.
func (app *App) saveFilteredPKToCSV() error {
	// Create name for CSV file.
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
	if errWriter := writer.Write([]string{app.Config.ColumnsOrder[0]}); errWriter != nil {
		return errWriter
	}

	// Write data body of the CSV file.
	for _, data := range app.Outputs.Data {
		// Take only PK column data.
		if errWriter := writer.Write([]string{data.ID}); errWriter != nil {
			return errWriter
		}
	}

	printStyled(
		fmt.Sprintf("– Saving filtered transactions to CSV file '%s' in the current dir... OK!", fileName),
		"margin-top",
	)

	return nil
}
