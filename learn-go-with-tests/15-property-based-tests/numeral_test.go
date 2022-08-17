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
