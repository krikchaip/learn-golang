package main

import "fmt"

func main() {
	sum := func(xs ...int) (result int) {
		fmt.Printf("%#v, ", xs)
		fmt.Printf("xs addr: %p, xs value addr(sum): %p\n", &xs, xs)

		for _, v := range xs {
			result += v
		}
		return
	}

	// ** `xs` in sum will be initialized with `nil`
	sum() // []int(nil), xs addr: 0x..., xs value addr(sum): 0x0

	// ** Go appends each argument in-order to the `xs` in `sum`
	sum(1, 2, 3, 4) // []int{1, 2, 3, 4}, xs addr: 0x..., xs value addr(sum): 0x...

	xs := []int{1, 2, 3, 4}
	fmt.Printf("xs addr: %p\n", &xs)            // xs addr: 0x...yy
	fmt.Printf("xs value addr(main): %p\n", xs) // xs value addr(main): 0x...a

	// ** Spread operator in Go behaves differently from Javascript
	// ** in a way that it passes `xs` as a slice header to `sum`
	// ** whenever you write `xs...`.
	sum(xs...) // []int{1, 2, 3, 4}, xs addr: 0x...zz, xs value addr(sum): 0x...a

	// ** unlike the spread operator in Javascript,
	// ** which expands each item in `xs` and pass it to `sum` as an argument
	// ** eg. sum(...xs) -> sum(xs[0], xs[1], ...)

	// ** this allows you to mutate the argument that you pass into. If it's a slice.
	mutate := func(arguments ...int) {
		arguments[0] = 1000
	}

	mutate(xs...)   // xs before: [1 2 3 4]
	fmt.Println(xs) // xs after: [1000 2 3 4]

	// ?? higher order function in go
	addN := func(n int) func(int) int {
		return func(x int) int {
			return x + n
		}
	}

	f := addN(10)
	fmt.Println(f(3)) // 13

	// ?? DEFER statement - postpones functions execution and will execute just before returns
	// ** execute in LIFO order (like stack)
	defer foo()
	bar()
	defer foobar()
	fmt.Println("Please ignore me...")

	// This is bar()!
	// Please ignore me...
	// This is foobar()!
	// This is foo()!
}

func foo() {
	fmt.Println("This is foo()!")
}

func bar() {
	fmt.Println("This is bar()!")
}

func foobar() {
	fmt.Println("This is foobar()!")
}
