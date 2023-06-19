package main

import (
	_ "embed"
	"fmt"
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

// removeDuplicates provides process to remove all duplicates in the given slice
// of slices, and return cleared slice ot slices.
func removeDuplicates(source [][]string) [][]string {
	// Create a temp map and slice.
	seen := make(map[string]bool, 0)
	result := make([][]string, 0)

	// Run loop for all slices in the given source.
	for _, item := range source {
		// Check, if current key is not in source slice.
		if _, ok := seen[fmt.Sprint(item)]; !ok {
			// Collect element into temp slice.
			result = append(result, item)

			// Set marker.
			seen[fmt.Sprint(item)] = true
		}
	}

	return result
}

// matchIndexes provides process to matching indexes for mapping columns.
func matchIndexes(wanted, source []string) map[string]int {
	// Create a temp map.
	indexes := make(map[string]int, 0)

	// Loop for all strings in the given wanted slice.
	for _, columnName := range wanted {
		// Loop for all strings in the given source slice.
		for sourceIndex, sourceName := range source {
			// Check, if names of the source and wanted slices were match.
			if sourceName == columnName {
				// Collect index of the source slice to the temp map.
				indexes[columnName] = sourceIndex
				break
			}
		}
	}

	return indexes
}

// matchValues provides process to matching two values with a condition.
func matchValues(value1, value2, condition string) (bool, error) {
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
				"wrong value '%s' to convert input value to type 'int' for this condition '%s'",
				value1, condition,
			)
		}

		// Convert value2 from string to int, or error.
		val2, err := strconv.Atoi(value2)
		if err != nil {
			return false, fmt.Errorf(
				"wrong value '%s' to convert filter value to type 'int' for this condition '%s'",
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

// printStyled provides a beauty output for console.
func printStyled(s, style string) {
	// Create a new blank style or the lipgloss.
	lp := lipgloss.NewStyle()

	// Switch between styles.
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
