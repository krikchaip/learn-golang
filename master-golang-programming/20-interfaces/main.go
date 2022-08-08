package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Object interface {
	Volume() float64
}

// ?? Embedded Interface - like TypeA & TypeB in Typescript :)
type Geometry interface {
	Shape  // includes every Shape's methods
	Object // includes every Object's methods
	GetColor() string
}

// implements: Shape
type Rectangle struct {
	Width, Height float64
}

// implements: Geometry
type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// ** Circle is the only struct that implements this method
func (c Circle) Volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.Radius, 3)
}

// ** also this method...
func (c Circle) GetColor() string {
	return "red"
}

func main() {
	s1 := Circle{Radius: 5}
	s2 := Rectangle{Width: 3, Height: 2}

	Print(s1) // Shape: main.Circle{Radius:5}, Area: 78.54, Perimeter: 31.42
	Print(s2) // Shape: main.Rectangle{Width:3, Height:2}, Area: 6.00, Perimeter: 10.00

	// ** uninitialized interface variable, default to <nil>
	var s Shape

	fmt.Printf("Type(s): %T\n", s)  // Type(s): <nil>
	fmt.Printf("Value(s): %v\n", s) // Value(s): <nil>

	Print(s) // nil Shape!

	// ** (Polymorphism) when assigned, s becomes a concrete value
	s = s1
	Print(s) // Shape: main.Circle{Radius:5}, Area: 78.54, Perimeter: 31.42

	// ** eventhough s is a Circle, we still cannot call s.Volume()
	// s.Volume()

	// ?? Type Assertion expressions
	s.(Circle).Volume()

	// ** this would fail because s is indeed a Circle
	// s.(Rectangle).Area()

	// ?? Type Assertion expressions - with ok status (similar to map values)
	if r, ok := s.(Rectangle); ok {
		fmt.Println("Rectangle Area:", r.Area())
	}

	// ?? Type Switches statement
	switch shape := s.(type) { // or switch s.(type) instead if you don't want to introduce a new variable
	case Rectangle:
		fmt.Println("shape is a Rectangle.")
	case Circle: // <- will match this case
		fmt.Println("shape is a Circle.")
		fmt.Println("its volume is equal to:", shape.Volume())
	}
}

// ?? This function accepts anything that implements Shape
func Print(s Shape) {
	// ?? It is possible to compare interface variables to nil
	if s == nil {
		fmt.Println("nil Shape!")
		return
	}

	fmt.Printf("Shape: %#v, ", s)
	fmt.Printf("Area: %.2f, ", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}
