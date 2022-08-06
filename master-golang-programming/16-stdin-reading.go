package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// ?? point a scanner to STDIN instead
	scanner := bufio.NewScanner(os.Stdin)

	for scanStdin(scanner) {
	}
}

func scanStdin(scanner *bufio.Scanner) bool {
	fmt.Print("Please enter something (type .quit to exit): ")

	ok := scanner.Scan()
	text := scanner.Text()

	if text == ".exit" {
		fmt.Println("\nGood Bye!")
		return false
	}

	fmt.Println("Input text:", scanner.Text())
	fmt.Println("Input bytes:", scanner.Bytes())
	fmt.Println()

	return ok
}
