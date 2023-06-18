package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"unicode/utf8"
)

// newInputs provides a new data input instance.
func newInputs(config *Config) (*Inputs, error) {
	// Check, if the input data file is not empty.
	if inputDataFilePath == "" {
		return nil, errors.New("file with input data is empty")
	}

	// Open the given file with input data.
	file, err := os.Open(filepath.Clean(inputDataFilePath))
	if err != nil {
		return nil, fmt.Errorf("can't open file with input data: %v", err)
	}
	defer file.Close()

	// Set up CSV reader.
	csvReader := csv.NewReader(bufio.NewReader(file))
	csvReader.LazyQuotes = true

	// Check, if the CSV column separator is not empty in config.
	if config.CSVSeparator != "" {
		// Decode rune from the string.
		separator, _ := utf8.DecodeRuneInString(config.CSVSeparator)

		// Set separator to the CSV reader.
		csvReader.Comma = separator
	}

	// Reading CSV.
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("can't read CSV file with input data: %v", err)
	}

	// Create a new Inputs struct.
	inputs := &Inputs{}

	// Loop for get mapping for fields of the CSV file.
	for index, data := range records {
		// Header is the first row.
		if index == 0 {
			// Collect header to mapping list.
			inputs.Mapping = matchIndexes(config.ColumnsOrder, data)
			continue
		}

		// Collect all other rows to data list.
		inputs.Data = append(inputs.Data, data)
	}

	// Remove duplicates.
	inputs.Data = removeDuplicates(inputs.Data)

	return inputs, nil
}
