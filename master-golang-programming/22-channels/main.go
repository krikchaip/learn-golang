package main

import (
	"fmt"
	"time"
)

func main() {
	// ?? declaring a bidirectional channel
	ch1 := make(chan int)

	fmt.Println(ch1) // 0x... (channel is a pointer)

	// ?? declaring and initilizing a RECEIVE-ONLY channel (unidirectional)
	ch2 := make(<-chan string)

	// ?? declaring and initilizing a SEND-ONLY channel (unidirectional)
	ch3 := make(chan<- string)

	fmt.Printf("%T, %T, %T\n", ch1, ch2, ch3) // chan int, <-chan string, chan<- string

	// ** using "send statement" inside the main goroutine will cause a DEADLOCK
	// ch1 <- 10

	// ** as opposed to the "receive expression" that is allowed
	// ** as long as there's some content inside
	// num := <-ch1

	// ** must call this inside a goroutine !!
	go putN(10, ch1)

	fmt.Println("Value received:", <-ch1) // 10

	ball := make(chan string)

	// ?? sending a message to "ball"
	go ping(ball)

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
	c <- n
}

func ping(c chan<- string) {
	c <- "PING"
}

func pong(c chan string) {
	c <- "PONG"
}
