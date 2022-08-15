package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Hello"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "World"
	}()

	// A select blocks until one of its cases can run, then it executes that case.
	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Printf("Received %q at %fs\n", msg, time.Since(start).Seconds())
		case msg := <-ch2:
			fmt.Printf("Received %q at %fs\n", msg, time.Since(start).Seconds())

			// Basic sends and receives on channels are blocking.
			// However, we can use `select` with a `default` clause to implement non-blocking channels.
			// default:
			// 	fmt.Println("noop...")
		}
	}

	total := time.Since(start).Seconds()

	fmt.Printf("total execution time: %fs\n", total) // ~2 seconds
}
