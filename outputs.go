package main

import (
	"fmt"

	"github.com/koddr/gosl"
)

// newOutputs provides a new output instance.
func newOutputs(config *Config, inputs *Inputs, filtered *Filtered) (*Outputs, error) {
	// Create a temp slice of slices for data.
	dataToOutput := inputs.Data

	// Check, if filtered data is not nil or empty.
	if filtered.Data != nil && len(filtered.Data) > 0 {
		// Set filtered data to the temp slice of slices.
		dataToOutput = filtered.Data
	}

	// Create a new Outputs instance.
	outputs := &Outputs{}

	// Loop for all data (into filtered or inputs slices).
	for _, data := range dataToOutput {
		// Loop for all update fields.
		for _, field := range config.UpdateFields {
			// Create a temp slice.
			m := make([]bool, 0)

			// Loop for checking condition.
			for index, condition := range field.Conditions {
				// Matching values with condition between input (in filtered or inputs) data
				// and column (in config file).
				match, err := matchValues(data[inputs.Mapping[condition.ColumnName]], condition.Value, condition.Condition)
				if err != nil {
					return nil, fmt.Errorf(
						"%w, column %d ('%v') in row %v",
						err, index, condition.ColumnName, data,
					)
				}

				// Collect matching result to the temp slice.
				m = append(m, match)
			}

			// Check, if there is no false in the conditions, it means
			// that the update fields settings have been respected.
			if !gosl.ContainsInSlice(m, false) {
				// Collect the current data to the outputs result.
				outputs.Data = append(outputs.Data, &Output{
					ID:        data[inputs.Mapping[config.ColumnsOrder[0]]],
					FieldName: field.FieldName,
					Values:    field.Values,
				})
			}
		}
	}

	return outputs, nil
}
