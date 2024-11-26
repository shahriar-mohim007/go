package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Start a goroutine that sends and receives values through channels
	go func() {
		v := 1
		ch1 <- v                         // Send value 1 to ch1
		v2 := <-ch2                      // Receive a value from ch2
		fmt.Println("Goroutine:", v, v2) // Print both values in the goroutine
	}()

	v := 2
	var v2 int
	select {
	case ch2 <- v:
	case v2 = <-ch1:

	}
	fmt.Println(v, v2)
}
