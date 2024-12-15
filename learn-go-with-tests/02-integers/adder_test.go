package main

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

// ?? run `godoc -http=:PORT` to see the doc
func ExampleAdd() {
	sum := Add(1, 5)

	// ** IMPORTANT: you must add the `Output: ...` comment
	// ** in order to make it run together with tests
	fmt.Println(sum) // Output: 6
}
