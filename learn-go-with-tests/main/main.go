package main

import (
	hello "01-hello"
	integers "02-integers"

	"fmt"
)

func main() {
	fmt.Println("01-hello:", hello.Hello("Winner", hello.French))
	fmt.Println("02-integers:", integers.Add(2, 5))
}
