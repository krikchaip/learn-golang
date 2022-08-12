package reflection_test

import (
	reflection "12-reflection"
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	x := struct{ Name string }{"Chris"}

	var got []string
	want := []string{x.Name}

	reflection.Walk(x, func(v string) {
		got = append(got, v)
	})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
