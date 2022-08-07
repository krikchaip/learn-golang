package main

import "fmt"

func main() {
	// ?? Constants need to be initialized when declared
	const days int = 7 // typed constant
	const pi = 3.14    // untyped constant

	// Declaring multiple (grouped) constants
	const (
		a         = 5   // untyped constant
		b float64 = 0.1 // typed constant
	)
	const n, m int = 4, 5

	const (
		min1 = -500
		max1 //gets its type and value form the previous constant. It's 500
		max2 //in a grouped constants, a constant repeats the previous one -> 500
	)

	// ** You cannot initiate a constant at runtime (constants belong to compile-time)
	// const power = math.Pow(2, 3) // error, functions calls belong to runtime

	// ** No Array or Struct constants
	// const y [2]int = {5, 6}
	// const y = [2]int{5, 6}

	// ** You cannot use a variable to initialize a constant
	// t := 5
	// const tc = t // error, variables belong to runtime and you cannot initialize a const to runtime values

	// ** You can use a function like len() to initialize a const if it has as argument
	const l1 = len("Hello") // OK
	// str := "Hello"; const l2 = len(str) // error, str is a variable and belongs to runtime

	const (
		x          = 5
		y  float64 = 1.1
		v1 int     = 5
		v2 float64 = 1.1
	)

	// => 5.5, No Error because x is untyped and gets its type when its used first time (float64).
	fmt.Println(x * y)

	// => Error: invalid operation: v1 * v2 (mismatched types int and float64)
	// fmt.Println(v1 * v2)

	// ?? IOTA -> a number generator for constants which starts from zero
	// ?? and is incremented by 1 automatically.
	const (
		c1 = iota // -> 0
		c2 = iota // -> 1
		c3 = iota // -> 2
	)

	const (
		North = iota // by default 0
		East         // omitting type and value means, repeating its type and value so East = iota = 1 (it increments by 1 automatically)
		South        // -> 2
		West         // -> 3
	)

	// Initializing the constants using a step:
	const (
		c11 = iota * 2 // -> 0
		c22            // -> 2
		c33            // -> 4
	)

	const (
		_         = iota
		Monday    // -> 1
		Tuesday   // -> 2
		Wednesday // -> 3
		Thursday  // -> 4
		Friday    // -> 5
		Saturday  // -> 6
		Sunday    // -> 7
	)
}
