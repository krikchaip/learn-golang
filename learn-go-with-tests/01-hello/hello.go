// to run this file, execute `go run hello.go`
// NOTE: this is only applicable to the `main` package
package main

import (
	"fmt"
)

func main() {
	fmt.Println("01-hello:", Hello("Winner", French))
}

// ?? UppercaseFunction -> public function
// alternative to function signature: `func Hello(name string, language string) string`
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

// ?? lowercaseFunction -> private function
// create a variable named `prefix`, assign it with the `zero` value
// and return at the end of the function.
func greetingPrefix(language string) (prefix string) {
	switch language {
	case French:
		prefix = frenchHelloPrefix
	case Spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

// ?? public global variables
const French = "French"
const Spanish = "Spanish"
const English = "Engligh"

// ?? private global variables
const frenchHelloPrefix = "Bonjour, "
const spanishHelloPrefix = "Hola, "
const englishHelloPrefix = "Hello, "
