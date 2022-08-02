package structs_methods_interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10., 10.}

	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	assert := func(t testing.TB, got, want float64) {
		t.Helper()

		// ?? use %g instead of %f when printing a higher-precision floating point number
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{Width: 12.0, Height: 6.0}

		got := rectangle.Area()
		want := 72.

		assert(t, got, want)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}

		got := circle.Area()
		want := 314.1592653589793

		assert(t, got, want)
	})

}
