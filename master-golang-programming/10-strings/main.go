package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// ?? a string literal enclosed in backticks is called a `raw string`,
	// ?? and it is interpreted literally.

	// ** backslashes or \n  have no special meaning
	fmt.Println(`Price: 100 \nBrand: Nike`)
	fmt.Println(`C:\Users\Andrei`)

	// escaping character in double-quotes
	fmt.Println("C:\\Users\\Andrei")

	// ?? get an element (byte) of a string
	s1 := "Winner"
	fmt.Println("Element at index 0:", s1[0]) // -> 87 (ascii code for 'W')

	// ?? a string is in fact a slice of bytes in Go
	var W byte = s1[0]
	fmt.Printf("Element at index 0: %c\n", W) // -> W

	str := []byte{97, 98, 99, 100, 101}
	fmt.Println(str)        // -> [97 98 99 100 101]
	fmt.Printf("%s\n", str) // -> abcde

	// ?? RUNE LITERALS - represents unicode points (int32)
	r1 := 'x'
	fmt.Printf("Type: %T, Value: %d\n", r1, r1)

	th := "เกริกชัย"
	fmt.Println(len(th))                          // -> 24bytes (3bytes for each Thai character)
	fmt.Println("Byte (not rune) at [1]:", th[1]) // -> Byte (not rune) at [1]: 185

	// ** WRONG - rune decoding in string
	for i := 0; i < len(th); i++ {
		fmt.Printf("%c", th[i]) // -> GIBBERISH 😡
	}

	fmt.Println()

	// ** BETTER - but not all runes are 3bytes
	for i := 0; i < len(th); i += 3 {
		fmt.Printf("%s", th[i:i+3]) // -> เกริกชัย 🤔
	}

	fmt.Println()

	// ** RIGHT - using the utf8 standard lib to handle each rune character size
	for i := 0; i < len(th); {
		r, size := utf8.DecodeRuneInString(th[i:])
		fmt.Printf("%c", r) // -> เกริกชัย 👍🏻
		i += size           // -> size = 3bytes in this case
	}

	fmt.Println()

	// ** BEST - just use the `for range` loop
	for _, c := range th {
		fmt.Printf("%c", c) // -> เกริกชัย 🤩
	}

	fmt.Println()

	// ?? counting the number of runes in a string
	fmt.Println(len("ţară"))                    // -> 6bytes (there are actually 4runes in the string)
	fmt.Println(utf8.RuneCountInString("ţară")) // -> 4

	// ** slicing a string returns []byte not []rune
	ss := "วินเนอร์ Krikchai!"
	fmt.Println(ss[2:5]) // -> GIBBERISH!!

	// ** you have to convert the string to []rune first before slicing
	rs := []rune(ss)
	fmt.Println(string(rs[:3])) // -> วิน
}
