package numeral

import (
	"strings"
)

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic uint16) string {
	// ?? A Builder is used to efficiently build a string using Write methods.
	// ?? It minimizes memory copying.
	var result strings.Builder // implements: io.Writer

	// ?? more efficient than string literal concatenation
	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) uint16 {
	var result uint16

	for i := 0; i < len(roman); {
		if i+1 < len(roman) && valueOf(roman[i]) < valueOf(roman[i+1]) {
			result += valueOf(roman[i+1]) - valueOf(roman[i])
			i += 2
			continue
		}

		result += valueOf(roman[i])
		i++
	}

	return result
}

func valueOf(c byte) uint16 {
	for _, numeral := range allRomanNumerals {
		if string(c) == numeral.Symbol {
			return numeral.Value
		}
	}
	return 0
}
