package numeral_test

import (
	numeral "15-property-based-tests"
	"fmt"
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	table := []struct {
		Arabic int
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
