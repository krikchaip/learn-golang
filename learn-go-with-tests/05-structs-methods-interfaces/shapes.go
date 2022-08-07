package structs_methods_interfaces

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.Width * rectangle.Height
}

type Circle struct {
	Radius float64
}

func (circle Circle) Area() float64 {
	return math.Pi * math.Pow(circle.Radius, 2)
}

// ?? short-form declaration
type Triangle struct {
	Base, Height float64
}

func (triangle Triangle) Area() float64 {
	return 0.5 * triangle.Base * triangle.Height
}
