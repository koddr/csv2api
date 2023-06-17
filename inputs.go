package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

// newInputs ...
func newInputs(config *Config) (*Inputs, error) {
	//
	file, err := os.Open(filepath.Clean(inputDataFilePath))
	if err != nil {
		return nil, fmt.Errorf("error openning CSV file with input data: %v", err)
	}
	defer file.Close()

	//
	csvReader := csv.NewReader(bufio.NewReader(file))
	csvReader.Comma = []rune(config.CSVSeparator)[0]
	csvReader.LazyQuotes = true

	//
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading CSV file with input data: %v", err)
	}

	//
	inputs := &Inputs{}

	//
	for index, data := range records {
		//
		if index == 0 {
			//
			inputs.Mapping = matchIndexes(config.ColumnsOrder, data)
			continue
		}

		//
		inputs.Data = append(inputs.Data, data)
	}

	//
	inputs.Data = removeDuplicates(inputs.Data)

	return inputs, nil
}
