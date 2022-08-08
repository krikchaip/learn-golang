package main

import "fmt"

type Car struct {
	brand string
	price int
}

// ** The changes are not visible outside because
// ** the receiver type is a struct, and it is passed BY VALUE.
func (c Car) NotMutate(brand string, price int) {
	c.brand = brand
	c.price = price
}

// ?? RECOMMENDED way to mutate a struct - Pointer Receiver
func (c *Car) Mutate(brand string, price int) {
	(*c).brand = brand // alias: c.brand = brand
	(*c).price = price // alias: c.price = price
}

func main() {
	carA := Car{"Audi", 40000}

	// ** nothing will happen...
	carA.NotMutate("Opel", 21000)
	fmt.Println(carA) // {Audi 40000}

	// ** mutate the struct directly
	(&carA).Mutate("Seat", 25000) // alias: carA.Mutate(...)
	fmt.Println(carA)             // {Seat 25000}

	// ** assign a pointer to struct in another variable and mutate
	carB := &carA
	fmt.Println(*carB) // {Seat 25000}

	carB.Mutate("VW", 30000) // alias: (&carA).Mutate(...)
	fmt.Println(*carB)       // {VW 30000}

	(*carB).Mutate("VW", 25000) // alias: carA.Mutate(...)
	fmt.Println(*carB)          // {VW 25000}

	// ** carA has also been modified
	fmt.Println(carA) // {VW 25000}
}

type DistancePtr *int

// ** method declarations are not permitted on named types
// ** that are themselves POINTER TYPES
// func (d DistancePtr) Invalid() {
// }
