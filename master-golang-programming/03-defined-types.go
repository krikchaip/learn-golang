package main

import "fmt"

// ?? `int` is the underlying type of a defined type `Speed`
type Speed int

func main() {
	s1 := Speed(3) // alternative: var s1 Speed = 3
	s2 := Speed(4) // alternative: var s2 Speed = 4
	fmt.Println(s1 - s2)

	var x int = 1

	// ** this will produce an error because they're different type
	// fmt.Println(s1 - x)

	// ** a type can be converted into another type if both types
	// ** share the same underlying type
	fmt.Println(s1 - Speed(x))
	fmt.Println(int(s1) - x)
	s1 = Speed(x)
	x = int(s1)

	type Km float32
	type Mile float32

	var parisToLondon Km = 465
	var inMiles Mile

	inMiles = Mile(parisToLondon) / 0.621
	fmt.Println(inMiles)
}
