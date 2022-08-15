package main

import (
	"fmt"
	"time"
)

func main() {
	// ?? declaring a bidirectional channel - unbuffered
	ch1 := make(chan int)

	fmt.Println(ch1) // 0x... (channel is a pointer)

	// ?? declaring and initilizing a RECEIVE-ONLY channel (unidirectional, unbuffered)
	ch2 := make(<-chan string)

	// ?? declaring and initilizing a SEND-ONLY channel (unidirectional, unbuffered)
	ch3 := make(chan<- string)

	fmt.Printf("%T, %T, %T\n", ch1, ch2, ch3) // chan int, <-chan string, chan<- string

	// ** calling "send statement" on an unbuffered channel
	// ** inside the main goroutine will result in a DEADLOCK
	// ch1 <- 10

	// ** as opposed to the "receive expression" that is always allowed.
	// ** it will block until there's some content inside a channel
	// num := <-ch1

	// ** must call this inside a goroutine (because ch1 is unbuffered) !!
	go putN(10, ch1)

	fmt.Println("Value received:", <-ch1) // 10

	// ?? buffered channel with capacity of 1
	ball := make(chan string, 1)

	fmt.Println("ball's capacity:", cap(ball)) // 1

	// ?? sending a message to "ball"
	ping(ball) // ** this will not block the main goroutine 🤩

	// ?? this is like calling "<-ball" continuously
	for res := range ball {
		fmt.Println(res)
		time.Sleep(1 * time.Second)
		switch res {
		case "PING":
			go pong(ball)
		case "PONG":
			go ping(ball)
		}
	}
}

func putN(n int, c chan int) {
	fmt.Println("putN started.")
	c <- n
	fmt.Println("putN finished.")
}

func ping(c chan<- string) {
	c <- "PING"
}

func pong(c chan string) {
	c <- "PONG"
}
