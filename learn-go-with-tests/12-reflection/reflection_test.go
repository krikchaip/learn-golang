package reflection_test

import (
	reflection "12-reflection"
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name   string
		X      any
		Result []string
	}{
		{
			Name:   "struct with one string field",
			X:      struct{ Name string }{"Chris"},
			Result: []string{"Chris"},
		},
		{
			Name:   "struct with multiple string fields",
			X:      struct{ Name, City string }{"Chris", "London"},
			Result: []string{"Chris", "London"},
		},
		{
			Name: "struct with non string field",
			X: struct {
				Name string
				Age  int
			}{"Chris", 33},
			Result: []string{"Chris"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			want := test.Result

			reflection.Walk(test.X, func(v string) {
				got = append(got, v)
			})

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}
}
