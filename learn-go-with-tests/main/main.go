package main

import (
	hello "01-hello"
	integers "02-integers"
	di "08-dependency-injection" // name is too long :(
	"fmt"
)

func main() {
	fmt.Println("01-hello:", hello.Hello("Winner", hello.French))
	fmt.Println("02-integers:", integers.Add(2, 5))
	di.GreetStdout()
	di.GreetHttp()
}
