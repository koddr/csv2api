package main

import (
	"errors"
	"fmt"

	"github.com/koddr/gosl"
)

// newFiltered provides a new filter instance.
func newFiltered(config *Config, inputs *Inputs) (*Filtered, error) {
	// Check, if the input data is not nil or empty.
	if inputs.Data == nil || len(inputs.Data) == 0 {
		return nil, errors.New("input data is empty or not parsed")
	}

	// Create a new Filtered instance.
	filtered := &Filtered{}

	// Check, if the filter has columns.
	if config.FilterColumns != nil && len(config.FilterColumns) > 0 {
		// Loop for all inputs data.
		for _, data := range inputs.Data {
			// Create a temp slice.
			m := make([]bool, 0)

			// Loop for all filtered columns.
			for index, column := range config.FilterColumns {
				// Matching values with condition between input (in data file) and column (in config file).
				match, err := matchValues(data[inputs.Mapping[column.ColumnName]], column.Value, column.Condition)
				if err != nil {
					return nil, fmt.Errorf(
						"%w, column %d ('%v') in row %v",
						err, index, column.ColumnName, data,
					)
				}

				// Collect matching result to the temp slice.
				m = append(m, match)
			}

			// Check, if there is no false in the conditions, it means
			// that the filter settings have been respected.
			if !gosl.ContainsInSlice(m, false) {
				// Collect the current data to the filtered result.
				filtered.Data = append(filtered.Data, data)
			}
		}
	}

	return filtered, nil
}
