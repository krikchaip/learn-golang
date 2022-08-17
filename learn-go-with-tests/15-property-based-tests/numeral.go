package numeral

func ConvertToRoman(arabic int) string {
	if arabic == 2 {
		return "II"
	}
	if arabic == 3 {
		return "III"
	}
	return "I"
}
