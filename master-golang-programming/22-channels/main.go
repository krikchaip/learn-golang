package main

import "fmt"

func main() {
	// ?? declaring a bidirectional channel
	ch1 := make(chan int)

	fmt.Println(ch1) // 0x... (channel is a pointer)

	// ?? Declaring and initilizing a RECEIVE-ONLY channel (unidirectional)
	ch2 := make(<-chan string)

	// ?? Declaring and initilizing a SEND-ONLY channel (unidirectional)
	ch3 := make(chan<- string)

	fmt.Printf("%T, %T, %T\n", ch1, ch2, ch3) // chan int, <-chan string, chan<- string

	// ** calling these statements alone inside main goroutine is not allowed!!
	// ch1 <- 10
	// num := <-ch1
	// _ = num

	// ** must call inside a goroutine !!
	go putN(10, ch1)
	n := <-ch1

	fmt.Println("Value received:", n) // 10
}

func putN(n int, c chan int) {
	c <- n
}
