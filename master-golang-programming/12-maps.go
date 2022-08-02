package main

import (
	"fmt"
)

type Tuple[T comparable] [2]T

func main() {
	// ?? the uninitialized (zero) value of maps is `nil` (just like slices)
	var employees map[string]string

	fmt.Printf("%#v\n", employees) // map[string]string(nil)
	fmt.Println(len(employees))    // 0

	// ?? if a key is not found, returns the zero value of value's type
	fmt.Printf("employees[\"Dan\"] = %q\n", employees["Dan"]) // employees["Dan"] = ""

	// ** like slices, you cannot assign a key-value pair on an uninitialized map
	// employees["Dan"] = "Programmer"

	// ** either assign a variable using an empty map literal
	// ** or calling `make` built-in function.
	employees = make(map[string]string) // alternative: employees = map[string]string{}
	employees["Dan"] = "Programmer"
	fmt.Println(employees) // map[Dan:Programmer]

	// ?? map keys can be any `comparable` (eg. boolean, string, int, float and array)
	m1 := map[Tuple[int]]string{{1, 1}: "Winner"}
	fmt.Println(m1[Tuple[int]{1, 1}]) // Winner

	// ?? check if a particular key EXISTS
	if _, exists := employees["Winner"]; !exists {
		fmt.Println(`"Winner" not found in employees`)
	}

	// ?? ITERATING over a map
	for k, v := range employees {
		fmt.Printf("Key: %#v, Value: %#v\n", k, v) // Key: "Dan", Value: "Programmer"
	}

	// ?? DELETING a key in a map
	delete(employees, "Dan")
	fmt.Println(employees) // map[]

	// ** When creating a map variable Go creates a pointer to a map header value in memory.
	// ** The key: value pairs of the map are not stored directly into the map.
	// ** They are stored in memory at the address referenced by the map header.

	friends := map[string]int{"Dan": 40, "Maria": 35}

	// ** neighbors shares the same reference as friends
	neighbors := friends
	friends["Dan"] = 30
	fmt.Println(neighbors) // map[Dan:30 Maria:35]

	// ?? COPYING map elements to other map
	colleagues := make(map[string]int)
	for k, v := range friends {
		colleagues[k] = v
	}
}
