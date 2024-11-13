package main

import (
	"fmt"
)

func main() {
	a := []int{2, 4, 6, 8, 10}
	ch := make(chan int, len(a)) // Buffered channel to avoid blocking

	for _, v := range a {
		go func() {
			ch <- v * 2
		}() // Pass v as an argument to avoid variable capture issues
	}

	// Retrieve and print results from channel
	for i := 0; i < len(a); i++ {
		fmt.Println(<-ch)
	}
}
