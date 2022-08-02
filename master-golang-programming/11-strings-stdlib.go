package main

import (
	"fmt"
	"strings"
)

func main() {
	// ?? substring CONTAINS
	fmt.Println(strings.Contains("I love Go!", "lovex")) // -> false

	// ?? CONTAINS ANY unicode code in chars
	fmt.Println(strings.ContainsAny("เกริกชัย", "เก")) // -> true

	// ?? CONTAINS a RUNE character
	fmt.Println(strings.ContainsRune("วินเนอร์", 'น')) // -> true

	// ?? COUNT substring
	fmt.Println(strings.Count("cheese", "e")) // -> 3
	fmt.Println(strings.Count("Five", ""))    // -> 5 (len(s) + 1)

	// ?? efficient case-ignoring string COMPARISON
	fmt.Println(strings.EqualFold("GO", "go")) // -> true

	// ** normally, we would compare 2 strings by using
	// **   strings.ToLower(s) == strings.ToLower(t)
	// ** but this isn't efficient at all because `strings.ToLower`
	// ** iterates over each rune one by one, converts it to lowercase
	// ** and returns the newly formed string before we compare each
	// ** string with the `==` operator.

	// ?? REPEATING a string pattern
	fmt.Println(strings.Repeat("Winner ", 3)) // -> Winner Winner Winner

	// ?? REPLACING string with a new string
	fmt.Println(strings.Replace("192.168.0.1", ".", "^", 2))  // -> 192^168^0.1
	fmt.Println(strings.Replace("192.168.0.1", ".", "^", -1)) // -> 192^168^0^1
	fmt.Println(strings.ReplaceAll("192.168.0.1", ".", "^"))  // -> 192^168^0^1

	// ?? SPLIT/JOIN a string by some seperator
	fmt.Println(strings.Split("a,b,c", ","))                // -> [a b c]
	fmt.Println(strings.Join([]string{"a", "b", "c"}, ",")) // -> a,b,c

	// ?? SPLIT a string by /\s+/ (white-spaces)
	fmt.Println(strings.Fields("Orange Green\nBlue	Yellow ")) // -> [Orange Green Blue Yellow]

	// ?? TRIMMING a string
	fmt.Println(strings.TrimSpace("\t Goodbye Windows, Welcome Linux!\n ")) // -> Goodbye Windows, Welcome Linux!
	fmt.Println(strings.Trim("...Hello, Gophers!!!?", ".!?"))               // -> Hello, Gophers
}
