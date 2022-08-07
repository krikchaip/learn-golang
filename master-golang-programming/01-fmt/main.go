package main

import "fmt"

func main() {
	// ?? fmt.Println() writes to standard output with a new line character
	fmt.Println("Hello Go World!")

	name, age := "Andrei", 35
	fmt.Println(name, "is", age, "years old.")

	// ?? fmt.Printf() prints out to stdout according to a format specifier called VERB
	a, b, c, grades := 10, 15.5, "Gophers", []int{10, 20, 30}
	const pi float64 = 3.14159265359
	student := Student{"Winner", 26}

	// %d -> decimal
	// %f -> float
	// %s -> string
	fmt.Printf("a is %d, b is %f, c is %s\n", a, b, c) // => a is 10, b is 15.500000, c is Gophers

	// %.4f -> formatting floats with 4 decimal points
	fmt.Printf("pi is %.4f\n", pi) // => pi is 3.1416

	// %q -> double-quoted string
	fmt.Printf("%q\n", c) // => "Gophers"

	// %v -> value (any)
	fmt.Printf("%v\n", grades)  // => [10 20 30]
	fmt.Printf("%v\n", student) // => {Winner 26}

	// %+v -> like %v but add struct field names
	fmt.Printf("%+v\n", student) // => {name:Winner age:26}

	// %#v -> a Go-syntax representation of the value
	fmt.Printf("%#v\n", grades)  // => []int{10, 20, 30}
	fmt.Printf("%#v\n", b)       // => 15.5
	fmt.Printf("%#v\n", student) // => main.Student{name:"Winner", age:26}

	// %T -> value Type
	fmt.Printf("b is %T. grades is %T\n", b, grades) // => b is float64. grades is []int

	// %t -> bool (true or false)
	fmt.Printf("Boolean value: %t\n", true) // => Boolean value: true

	// %p -> pointer (address in base 16, with leading 0x)
	fmt.Printf("The address of a: %p\n", &a) // => The address of a: 0xSOMETHING

	// %c -> char (rune) represented by the corresponding Unicode code point
	fmt.Printf("%c and %c\n", 100, 51011) // => d and 읃  (runes for code points 101 and 51011)

	// %b -> base 2
	fmt.Printf("255 in base 2 is %b\n", 255) // => 255 in base 2 is 11111111

	// %08b -> base 2 with 8 digits
	fmt.Printf("32 in base 2 is %08b\n", 32) // => 32 in base 2 is 00100000

	// %x -> base 16
	fmt.Printf("101 in base 16 is %x\n", 101) // => 101 in base 16 is 65

	// ?? fmt.Sprintf() returns a string. Uses the same verbs as fmt.Printf()
	s := fmt.Sprintf("a is %d, b is %f, c is %s \n", a, b, c)
	fmt.Println(s) // a is 10, b is 15.500000, c is Gophers
}

type Student struct {
	name string
	age  int
}
