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
