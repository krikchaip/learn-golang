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
		{Name: "simple", X: struct{ Name string }{"Chris"}, Result: []string{"Chris"}},
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
