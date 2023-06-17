package main

import (
	"errors"
	"fmt"

	"github.com/koddr/gosl"
)

// newOutputs ...
func newOutputs(config *Config, inputs *Inputs, filtered *Filtered) (*Outputs, error) {
	//
	if filtered.Data == nil || len(filtered.Data) == 0 {
		return nil, errors.New("error: file with input data is empty")
	}

	//
	outputs := &Outputs{}

	//
	for _, data := range filtered.Data {
		//
		for _, field := range config.UpdateFields {
			//
			m := make([]bool, 0)

			//
			for index, condition := range field.Conditions {
				//
				inputValue := fmt.Sprint(data[inputs.Mapping[condition.ColumnName]])

				//
				match, err := matchValues(inputValue, condition.Value, condition.Condition)
				if err != nil {
					return nil, fmt.Errorf(
						"%w, column %d ('%v') in row %v",
						err, index, condition.ColumnName, data,
					)
				}

				//
				m = append(m, match)
			}

			//
			if !gosl.ContainsInSlice(m, false) {
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
