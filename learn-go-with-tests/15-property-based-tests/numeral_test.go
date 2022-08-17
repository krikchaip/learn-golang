package numeral_test

import (
	numeral "15-property-based-tests"
	"fmt"
	"testing"
	"testing/quick"
)

var table = []struct {
	Arabic uint16
	Roman  string
}{
	{Arabic: 1, Roman: "I"},
	{Arabic: 2, Roman: "II"},
	{Arabic: 3, Roman: "III"},
	{Arabic: 4, Roman: "IV"},
	{Arabic: 5, Roman: "V"},
	{Arabic: 6, Roman: "VI"},
	{Arabic: 7, Roman: "VII"},
	{Arabic: 9, Roman: "IX"},
	{Arabic: 10, Roman: "X"},
	{Arabic: 14, Roman: "XIV"},
	{Arabic: 18, Roman: "XVIII"},
	{Arabic: 20, Roman: "XX"},
	{Arabic: 39, Roman: "XXXIX"},
	{Arabic: 40, Roman: "XL"},
	{Arabic: 47, Roman: "XLVII"},
	{Arabic: 49, Roman: "XLIX"},
	{Arabic: 50, Roman: "L"},
	{Arabic: 90, Roman: "XC"},
	{Arabic: 100, Roman: "C"},
	{Arabic: 400, Roman: "CD"},
	{Arabic: 500, Roman: "D"},
	{Arabic: 798, Roman: "DCCXCVIII"},
	{Arabic: 900, Roman: "CM"},
	{Arabic: 1000, Roman: "M"},
	{Arabic: 1006, Roman: "MVI"},
	{Arabic: 1984, Roman: "MCMLXXXIV"},
	{Arabic: 1996, Roman: "MCMXCVI"},
	{Arabic: 2014, Roman: "MMXIV"},
	{Arabic: 2539, Roman: "MMDXXXIX"},
	{Arabic: 3999, Roman: "MMMCMXCIX"},
}

func TestConvertToRoman(t *testing.T) {
	for _, test := range table {
		name := fmt.Sprintf("%d->%s", test.Arabic, test.Roman)
		t.Run(name, func(t *testing.T) {
			got := numeral.ConvertToRoman(test.Arabic)
			want := test.Roman

			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})
	}
}

func TestConvertToArabic(t *testing.T) {
	for _, test := range table {
		name := fmt.Sprintf("%s->%d", test.Roman, test.Arabic)
		t.Run(name, func(t *testing.T) {
			got := numeral.ConvertToArabic(test.Roman)
			want := test.Arabic

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}

// ?? property-based testing
func TestPropertyOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		// ?? property exclusion
		if arabic < 1 || arabic > 3999 {
			return true
		}

		roman := numeral.ConvertToRoman(arabic)
		fromRoman := numeral.ConvertToArabic(roman)
		return arabic == fromRoman
	}

	// ?? 1000 samples
	config := &quick.Config{MaxCount: 1000}

	if err := quick.Check(assertion, config); err != nil {
		t.Error("failed checks", err)
	}
}
