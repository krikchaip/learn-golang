package numeral_test

import (
	numeral "15-property-based-tests"
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	table := []struct {
		name   string
		input  int
		output string
	}{
		{name: "1->I", input: 1, output: "I"},
		{name: "2->II", input: 2, output: "II"},
		{name: "3->III", input: 3, output: "III"},
		{name: "4->IV", input: 4, output: "IV"},
		{name: "5->V", input: 5, output: "V"},
		{name: "6->VI", input: 6, output: "VI"},
		{name: "7->VI", input: 7, output: "VII"},
		{name: "9->IX", input: 9, output: "IX"},
		{name: "10->X", input: 10, output: "X"},
		{name: "14->XIV", input: 14, output: "XIV"},
		{name: "18->XVIII", input: 18, output: "XVIII"},
		{name: "20->XX", input: 20, output: "XX"},
		{name: "39->XXXIX", input: 39, output: "XXXIX"},
	}

	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			got := numeral.ConvertToRoman(test.input)
			want := test.output

			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})
	}
}
