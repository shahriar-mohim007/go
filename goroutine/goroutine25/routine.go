package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout! No message received.")
	}
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() { ch1 <- "message from channel 1" }()
	go func() { ch2 <- "message from channel 2" }()

	select {
	case msg := <-ch1:
		fmt.Println("Received:", msg)
	case msg := <-ch2:
		fmt.Println("Received:", msg)
	}

	select {
	case ch <- "hello":
		fmt.Println("Sent message")
	default:
		fmt.Println("Channel not ready, skipping send")
	}
	ch3 := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch3 <- i
		}
		close(ch3)
	}()
	for msg := range ch3 {
		fmt.Println("Received:", msg)
	}
	for {
		select {
		case val, ok := <-ch3:
			if !ok {
				fmt.Println("Channel closed")
				return
			}
			fmt.Println("Received:", val)
		}
	}

}
