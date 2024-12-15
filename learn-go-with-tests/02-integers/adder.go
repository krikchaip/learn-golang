package main

import "fmt"

func main() {
	fmt.Println("02-integers:", Add(2, 5))
}

// Add takes two integers and returns the sum of them.
func Add(x, y int) int {
	return x + y
}
