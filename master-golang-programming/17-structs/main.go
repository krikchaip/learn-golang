package main

import "fmt"

type Book struct {
	title       string
	author      string
	year, pages int // ?? multiple declaration
}

type Author struct {
	name  string
	books []Book
}

func main() {
	// ** this is a struct literal and ORDER MATTERS
	book_a := Book{"The Divine Comedy", "Dante Aligheri", 1320, 100}

	// ** declare by specifying fields -> ORDER DOESN'T MATTER
	book_b := Book{year: 1995, title: "Animal Farm", pages: 300, author: "George Orwell"}

	// ** omitting some fields -> will initialize with the zero-value
	book_c := Book{title: "Just a random book"}
	fmt.Printf("%#v\n", book_c) // main.Book{title:"Just a random book", author:"", year:0, pages:0}

	fmt.Println(book_a)         // {The Divine Comedy Dante Aligheri 1320 100}
	fmt.Printf("%v\n", book_a)  // {The Divine Comedy Dante Aligheri 1320 100}
	fmt.Printf("%#v\n", book_a) // main.Book{title:"The Divine Comedy", author:"Dante Aligheri", year:1320, pages:100}

	// ?? %+v -> like %v but add struct field names
	fmt.Printf("%+v\n", book_a) // {title:The Divine Comedy author:Dante Aligheri year:1320 pages:100}

	// ?? COMPARING struct values
	// ?? they are equal if their corresponding fields are equal.
	fmt.Println(book_a == book_b) // false

	john := Author{"John", []Book{book_a}}
	jane := Author{"Jane", []Book{book_a}}

	// ** cannot compare directly, because Author struct contains a slice field
	// fmt.Println(john == jane)

	_, _ = john, jane

	// ?? structs are like arrays, their value is atomic
	book_d := book_a // book_d is an ACTUAL COPY of book_a with different memory location
	book_d.title = "BULLSHIT!"

	fmt.Printf("%+v\n", book_a) // {title:The Divine Comedy author:Dante Aligheri year:1320 pages:100}
	fmt.Printf("%+v\n", book_d) // {title:BULLSHIT! author:Dante Aligheri year:1320 pages:100}
}
