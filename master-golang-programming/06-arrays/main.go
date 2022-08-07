package main

import "fmt"

func main() {
	// ?? declaring an array with four zero-values
	var numbers [4]int
	var floats = [4]float64{}

	// array zero value is zeroed value elements
	fmt.Printf("%v\n", numbers)  // -> [0 0 0 0]
	fmt.Printf("%#v\n", numbers) // -> [4]int{0, 0, 0, 0}
	fmt.Printf("%v\n", floats)   // -> [0 0 0 0]
	fmt.Printf("%#v\n", floats)  // -> [4]float64{0, 0, 0, 0}

	// ?? initializing only the first 2 elements
	ss := [4]string{"x", "y"}
	fmt.Printf("%v\n", ss) // -> [x y  ]

	// ?? the ellipsis operator (...) finds out automatically the length of the array
	len3 := [...]int{1, 4, 5}
	len3multiline := [...]int{
		1,
		4,
		5,
	}

	fmt.Println(len(len3), len(len3multiline)) // -> 3 3

	// ?? multi-dimensional array
	balances := [2][3]int{
		[3]int{5, 6, 7}, // [3]int is optional
		{8, 9, 10},
	}

	fmt.Println(balances)

	// ** these arrays are not connected and are saved in different memory locations
	m := [3]int{1, 2, 3}
	n := m // n is an ACTUAL COPY of m

	fmt.Println("n is equal to m:", n == m) // => true
	m[0] = -1                               // only m is modified
	fmt.Println("n is equal to m:", n == m) // => false

	// ?? Arrays with Keyed Elements

	//the keyed elements can be in any order
	// -> [5 10 7]
	fmt.Println([3]int{
		1: 10,
		0: 5,
		2: 7,
	})

	// -> [0 0 50]
	fmt.Println([3]int{2: 50})

	// -> 5
	fmt.Println(len([...]string{4: "Dan"}))

	// ** un unkeyed element gets its index from the last keyed element
	// -> [7]string{"", "NYC", "Bangkok", "", "", "Paris", "London"}
	fmt.Printf("%#v\n", [...]string{
		5:         "Paris",
		"London",  // this is at index 6
		1:         "NYC",
		"Bangkok", // this is at index 2
	})
}
