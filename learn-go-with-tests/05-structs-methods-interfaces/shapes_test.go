package structs_methods_interfaces

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10., 10.}

	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	assert := func(t testing.TB, shape Shape, want float64) {
		t.Helper()

		// ?? calling interface method
		got := shape.Area()

		// ?? use %g instead of %f when printing a higher-precision floating point number
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	// ?? table driven tests
	areaTests := []struct {
		shape   Shape
		hasArea float64
	}{
		{Rectangle{12.0, 6.0}, 72.0},                          // short-form
		{Circle{Radius: 10.0}, 314.1592653589793},             // initialze struct explicitly
		{shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0}, // all combined
	}

	for _, row := range areaTests {
		shape := row.shape
		want := row.hasArea

		testName :=
			fmt.Sprintf("%s%v", reflect.TypeOf(shape).Name(), shape) +
				" -> " +
				fmt.Sprintf("%g", want)

		t.Run(testName, func(t *testing.T) {
			assert(t, shape, want)
		})
	}
}
