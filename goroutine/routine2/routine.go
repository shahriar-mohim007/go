package main

import "fmt"

func main() {
	ch1 := make(chan int, 2)
	ch2 := make(chan int, 2)
	//ch1 := make(chan int)
	//ch2 := make(chan int)
	// deadlocking goroutine

	// Start a goroutine that sends and receives values through channels
	go func() {
		v := 1
		ch1 <- v                         // Send value 1 to ch1
		v2 := <-ch2                      // Receive a value from ch2
		fmt.Println("Goroutine:", v, v2) // Print both values in the goroutine
	}()

	v := 2
	ch2 <- v                    // Send value 2 to ch2
	v2 := <-ch1                 // Receive a value from ch1
	fmt.Println("Main:", v, v2) // Print both values in the main function
}

//Diagram After Step 1:

//+---------------------+                     +------------------+
//|  Anonymous Goroutine|                     | Main Goroutine   |
//|---------------------|                     |------------------|
//| ch1 <- 1 (BLOCKED)  |                     |                  |
//| Wait to receive ch2 |                     |                  |
//+---------------------+                     +------------------+
//|                                        |
//V                                        V
//ch1 <------------------------------------> ch2

//Step 2: Main Goroutine Sends to ch2

//+---------------------+                     +---------------------+
//|  Anonymous Goroutine|                     |  Main Goroutine     |
//|---------------------|                     |---------------------|
//| ch1 <- 1 (BLOCKED)  |  <-- Blocked on --> | ch2 <- 2 (BLOCKED)  |
//| Wait to receive ch2 |                     | Wait to receive ch1 |
//+---------------------+                     +---------------------+
//|                                        |
//V                                        V
//ch1 <------------------------------------> ch2
