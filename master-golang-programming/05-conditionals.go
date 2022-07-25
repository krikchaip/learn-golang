package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// ?? IF statement variants

	// ** i and err are variables scoped to the if statement only
	if i, err := strconv.Atoi("34"); err == nil {
		fmt.Println("No error. i is", i)
	} else {
		fmt.Println(err)
	}

	// ** multiple conditions - variables are shared across ifs/else
	if args := os.Args; len(args) != 2 {
		fmt.Println("One argument is required!")
	} else if km, err := strconv.Atoi(args[1]); err != nil {
		fmt.Println("The argument must be an integer! Error:", err)
	} else {
		fmt.Printf("%d km in miles is %.4f\n", km, float64(km)/1.609)
	}
}
