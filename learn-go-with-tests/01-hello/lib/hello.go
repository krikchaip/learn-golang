package hello

// ?? UppercaseFunction -> public function
// alternative: `Hello(name string, language string)`
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
