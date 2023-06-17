package main

import (
	_ "embed"
	"fmt"
	"reflect"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/koddr/gosl"
)

const (
	// WikiPageURL link to project's Wiki page.
	WikiPageURL string = "https://github.com/koddr/csv2api/wiki"
)

var (
	//go:embed embed/config.yml
	embedConfigYAMLFile []byte

	//go:embed embed/data.csv
	embedDataCSVFile []byte

	// Flag for generate init config.
	initConfig bool

	// Flags for config, data file path, env prefix.
	configFilePath, inputDataFilePath, envPrefix string
)

// removeDuplicates ...
func removeDuplicates(source [][]string) [][]string {
	//
	seen := make(map[string]bool, 0)
	result := make([][]string, 0)

	//
	for _, item := range source {
		//
		if _, ok := seen[fmt.Sprint(item)]; !ok {
			//
			result = append(result, item)

			//
			seen[fmt.Sprint(item)] = true
		}
	}

	return result
}

// matchIndexes ...
func matchIndexes(wanted, source []string) map[string]int {
	//
	indexes := make(map[string]int, 0)

	//
	for _, columnName := range wanted {
		//
		for sourceIndex, sourceName := range source {
			//
			if sourceName == columnName {
				//
				indexes[columnName] = sourceIndex
				break
			}
		}
	}

	return indexes
}

// matchValues ...
func matchValues(value1, value2 string, condition string) (bool, error) {
	// Check condition.
	switch condition {
	case "EQ":
		// Equals to the given value (==).
		return value1 == value2, nil
	case "NEQ":
		// Not equal to the given value (!=).
		return value1 != value2, nil
	case "GT", "LT", "GTE", "LTE":
		// Convert value1 from string to int, or error.
		val1, err := strconv.Atoi(value1)
		if err != nil {
			return false, fmt.Errorf(
				"error: wrong value '%s' to convert input value to type 'int' for this condition '%s'",
				value1, condition,
			)
		}

		// Convert value2 from string to int, or error.
		val2, err := strconv.Atoi(value2)
		if err != nil {
			return false, fmt.Errorf(
				"error: wrong value '%s' to convert filter value to type 'int' for this condition '%s'",
				value2, condition,
			)
		}

		// Check condition.
		switch condition {
		case "GT":
			// Greater than the given value (>).
			return val1 > val2, nil
		case "LT":
			// Less than the given value (<).
			return val1 < val2, nil
		case "GTE":
			// Greater than or equal to the given value (>=).
			return val1 >= val2, nil
		case "LTE":
			// Less than or equal to the given value (<=).
			return val1 <= val2, nil
		}
	}

	return false, fmt.Errorf("unknown condition: %s", condition)
}

// ModifyByValue ...
func ModifyByValue(m map[string]any, foundValue, newValue any) (bool, map[string]any) {
	//
	var foundKey bool

	//
	if m == nil {
		return foundKey, nil
	}

	//
	for k, v := range m {
		if reflect.DeepEqual(v, foundValue) {
			//
			m[k] = newValue
			foundKey = true
		} else if mv, ok := v.(map[string]any); ok {
			//
			isFound, mod := ModifyByValue(mv, foundValue, newValue)

			//
			if isFound {
				m[k] = mod
				foundKey = true
			}
		}
	}

	return foundKey, m
}

// printStyled ...
func printStyled(s, style string) {
	lp := lipgloss.NewStyle()

	switch style {
	case "margin-top-bottom":
		fmt.Println(gosl.RenderStyled(s, lp.MarginTop(1).MarginBottom(1)))
	case "margin-top":
		fmt.Println(gosl.RenderStyled(s, lp.MarginTop(1)))
	case "margin-bottom":
		fmt.Println(gosl.RenderStyled(s, lp.MarginBottom(1)))
	case "margin-left":
		fmt.Println(gosl.RenderStyled(s, lp.MarginLeft(1)))
	default:
		fmt.Println(s)
	}
}
