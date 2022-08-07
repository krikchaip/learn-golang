package main

import "fmt"

func main() {
	// ?? APPEND slices (behave like list.concat() in other language)
	numbers := []int{2, 3}

	fmt.Println(append(numbers, 4)) // => [2 3 4]
	fmt.Println(numbers)            // => [2 3]

	// ** you have to reassign the variable to make it work
	numbers = append(numbers, 4)
	fmt.Println(numbers) // => [2 3 4]

	// ?? APPEND multiple values
	fmt.Println(append(numbers, 5, 6, 7)) // => [2 3 4 5 6 7]

	// ** spread operator for slices (like in Javascript/Python)
	n := []int{100, 200, 300}
	fmt.Println(append(numbers, n...)) // => [2 3 4 100 200 300]

	// ?? COPY slices (in-place operation)
	src := []int{10, 20, 30}
	dst := make([]int, len(src)) // => [0 0 0]

	nCopied := copy(dst, src)
	fmt.Println("copied", nCopied, "elements") // => copied 3 elements
	fmt.Println("dst:", dst)                   // => dst: [10 20 30]

	// ** dst slice has a shorter length than src
	dst = make([]int, 1)

	nCopied = copy(dst, src)
	fmt.Println("copied", nCopied, "elements") // => copied 1 elements
	fmt.Println("dst:", dst)                   // => dst: [10]

	// ** zero size dst slice -> noop
	dst = make([]int, 0)

	nCopied = copy(dst, src)
	fmt.Println("copied", nCopied, "elements") // => copied - elements
	fmt.Println("dst:", dst)                   // => dst: []
}
