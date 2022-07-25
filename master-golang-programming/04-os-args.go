package main

import (
	"fmt"
	"os"
)

func main() {
	// ** the first argument is always the path to the program
	// => Output: "os.Args([]string): [/path/to/program.go, ...]"
	fmt.Printf("os.Args(%T): %v\n", os.Args, os.Args)

	// ** the argument type is always a string so you need to parse it
	// ** before continue further
	fmt.Println("\n[total arguments:", fmt.Sprintf("%d", len(os.Args))+"]")
	for i := 0; i < len(os.Args); i++ {
		fmt.Printf("argument[%d, %T]: %v\n", i, os.Args[i], os.Args[i])
	}
}
