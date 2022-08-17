package numeral

import (
	"strings"
)

func ConvertToRoman(arabic int) string {
	// ?? A Builder is used to efficiently build a string using Write methods.
	// ?? It minimizes memory copying.
	var result strings.Builder // implements: io.Writer

	// ?? more efficient than string literal concatenation
	for arabic > 0 {
		switch {
		case arabic > 9:
			result.WriteString("X")
			arabic -= 10
		case arabic > 8:
			result.WriteString("IX")
			arabic -= 9
		case arabic > 4:
			result.WriteString("V")
			arabic -= 5
		case arabic > 3:
			result.WriteString("IV")
			arabic -= 4
		default:
			result.WriteString("I")
			arabic--
		}
	}

	return result.String()
}
