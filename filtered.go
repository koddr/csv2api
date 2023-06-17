package main

import (
	"errors"
	"fmt"
	"github.com/koddr/gosl"
)

// newFiltered ...
func newFiltered(config *Config, inputs *Inputs) (*Filtered, error) {
	//
	if inputs.Data == nil || len(inputs.Data) == 0 {
		return nil, errors.New("error: file with input data is empty")
	}

	//
	filtered := &Filtered{}

	//
	if config.FilterColumns != nil {
		//
		if len(config.FilterColumns) > 0 {
			//
			for _, data := range inputs.Data {
				//
				m := make([]bool, 0)

				//
				for index, column := range config.FilterColumns {
					//
					inputValue := fmt.Sprint(data[inputs.Mapping[column.ColumnName]])

					//
					match, err := matchValues(inputValue, column.Value, column.Condition)
					if err != nil {
						return nil, fmt.Errorf(
							"%w, column %d ('%v') in row %v",
							err, index, column.ColumnName, data,
						)
					}

					//
					m = append(m, match)
				}

				//
				if !gosl.ContainsInSlice(m, false) {
					//
					filtered.Data = append(filtered.Data, data)
				}
			}
		}
	}

	return filtered, nil
}
