package utils

import (
	"fmt"
	"strings"
)

func PrinterOneDimensionalSlice[T any](slice []T) string {
	var result []string
	result = append(result, fmt.Sprintf("n = %v\n  ", len(slice)))

	for _, val := range slice {
		result = append(result, fmt.Sprint(val))
	}

	return strings.Join(result, "  ")
}

func PrinterTwoDimensionalSlice[T any](slice [][]T) string {
	var result []string
	result = append(result, fmt.Sprintf("n = %v", len(slice)))

	maxWidth := 0
	for _, row := range slice {
		for _, cell := range row {
			s := fmt.Sprint(cell)
			if len(s) > maxWidth {
				maxWidth = len(s)
			}
		}
	}

	for _, row := range slice {
		var rowResult []string
		for _, cell := range row {
			rowResult = append(rowResult, fmt.Sprintf("%*s", maxWidth, fmt.Sprint(cell)))
		}
		result = append(result, strings.Join(rowResult, "  "))
	}

	return strings.Join(result, "\n")
}
