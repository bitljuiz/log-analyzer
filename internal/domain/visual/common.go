package visual

import (
	"fmt"
	"strings"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/models/analyzer"
)

var commonInformationOrder = []string{
	"File(-s)",
	"From data",
	"To data",
	"Requests count",
	"Minimum request size",
	"Maximum request size",
	"Average request size",
	"Percentile",
}

// FormatWithUnderscores форматирует строку числа, добавляя символ `_` как разделитель тысяч.
// FormatWithUnderscores("1000000") = "1_000_000".
func FormatWithUnderscores(n string) string {
	if len(n) <= 3 {
		return n
	}

	reversed := Reverse(n)

	var parts []string

	for i := 0; i < len(reversed); i += 3 {
		end := i + 3
		if end > len(reversed) {
			end = len(reversed)
		}

		parts = append(parts, reversed[i:end])
	}

	return Reverse(strings.Join(parts, "_"))
}

// FormatFilenames форматирует список имён файлов для отображения.
// FormatFilenames([]string{"1", "2"}) = "`1` `2` ".
func FormatFilenames(files []string) string {
	sb := strings.Builder{}

	for _, file := range files {
		sb.WriteString(fmt.Sprintf("`%s` ", file))
	}

	return sb.String()
}

// Reverse разворачивает строку.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

// OutputToCommon создаёт мапу значений для общей информации о статистике и форматирует данные.
func OutputToCommon(data *analyzer.Statistics) CommonInformation {
	return CommonInformation{
		"File(-s)":             FormatFilenames(data.Files),
		"From data":            data.From,
		"To data":              data.To,
		"Requests count":       FormatWithUnderscores(data.TotalRequestsNumber.String()),
		"Minimum request size": FormatWithUnderscores(fmt.Sprintf("%d", data.MinSizeRequest)) + "b",
		"Maximum request size": FormatWithUnderscores(fmt.Sprintf("%d", data.MaxSizeRequest)) + "b",
		"Average request size": FormatWithUnderscores(data.AverageRequestNumber.String()) + "b",
		"Percentile":           FormatWithUnderscores(fmt.Sprintf("%d", data.Percentile)) + "b",
	}
}
