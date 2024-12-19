package main

import (
	"fmt"
	"sync"
)

func fanOut(in <-chan int, workerCount int) []<-chan int {
	outs := make([]<-chan int, workerCount)
	for i := 0; i < workerCount; i++ {
		out := make(chan int)
		outs[i] = out
		go func(ch chan int) {
			for num := range in {
				ch <- num * num // Simulate processing
			}
			close(ch)
		}(out)
	}
	return outs
}

func fanIn(channels []<-chan int, out chan<- int) {
	var wg sync.WaitGroup
	wg.Add(len(channels))

	for _, ch := range channels {
		go func(c <-chan int) {
			for val := range c {
				out <- val
			}
			wg.Done()
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
}

func main() {
	input := make(chan int)
	output := make(chan int)

	// Fan-Out to multiple workers
	outs := fanOut(input, 3)

	// Fan-In from multiple workers to one channel
	fanIn(outs, output)

	go func() {
		for i := 1; i <= 10; i++ {
			input <- i
		}
		close(input)
	}()

	for result := range output {
		fmt.Println("Result:", result)
	}
}
