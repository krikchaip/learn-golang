package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
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

	// ?? FOR loop variants (there's only FOR loop in go. No WHILE, DO-WHILE)

	for i := 0; i < 10; i++ {
		// do something
	}

	// ** the last term can be omitted
	for i := 0; i < 10; {
		// do something
		i++
	}

	// ** WHILE loop
	i := 3
	cond := func() bool { return false }

	for i >= 0 {
		// do something
		i--
	}

	for cond() {
		// do something
	}

	// ** infinite loop (FOR without conditions)
	expo2 := 2
	for {
		expo2 *= 2
		if expo2 >= 1024 {
			fmt.Println(expo2)
			break
		}
	}

	// ** iterating over an array/slice
	xs := []string{"W", "I", "N", "N", "E", "R"}
	for i, char := range xs {
		fmt.Printf("i: %v, char: %v\n", i, char)
	}

	// ?? SWITCH case variants

	// ** no need for BREAK statements
	language := "Golang"
	switch language {
	case "Python":
		fmt.Println("Python")
	case "Go", "Golang": // "Go" || "Golang"
		fmt.Println("Golang")
	default: // default case is optional (if not match any case then do nothing)
		fmt.Println("Anything")
	}

	// ** cleaner IFs statements using SWITCH statement 👍🏻
	hour := time.Now().Hour()
	switch { // alternative: switch true { ... }
	case hour < 12:
		fmt.Println("Good Morning!")
	case hour < 17:
		fmt.Println("Good Afternoon!")
	default:
		fmt.Println("Good Evening!")
	}

	// ** even shorter syntax
	switch n := 10; {
	case n > 0:
		fmt.Println("Positive")
	case n < 0:
		fmt.Println("Negative")
	default:
		fmt.Println("Zero")
	}
}
