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
	rectangle := Rectangle{Width: 12.0, Height: 6.0}

	got := Area(rectangle)
	want := 72.

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
