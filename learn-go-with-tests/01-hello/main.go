package main

import (
	hello "01-hello/lib"
	"fmt"
)

func main() {
	fmt.Println("01-hello:", hello.Hello("Winner", hello.French))
}
