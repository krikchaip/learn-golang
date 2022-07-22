package main

import (
	hello "01-hello"
	integers "02-integers"
	"fmt"
)

func main() {
	fmt.Println(hello.Hello("world", ""))
	fmt.Println(integers.Add(2, 5))
}
