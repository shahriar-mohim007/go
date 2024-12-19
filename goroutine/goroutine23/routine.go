package main

import "fmt"

// When multiple goroutines send data to a single channel or data
// from one channel is distributed across multiple channels, select can help manage these patterns.
func fanOut(in <-chan int, workerCount int) {
	for i := 0; i < workerCount; i++ {
		go func(workerID int) {
			for num := range in {
				fmt.Printf("Worker %d processed %d\n", workerID, num)
			}
		}(i)
	}
}

func fanIn(ch1, ch2 <-chan string, out chan<- string) {
	go func() {
		for {
			select {
			case msg, ok := <-ch1: // Check if ch1 is closed
				if ok {
					out <- msg
				} else {
					ch1 = nil // Prevent further checks on ch1
				}
			case msg, ok := <-ch2: // Check if ch2 is closed
				if ok {
					out <- msg
				} else {
					ch2 = nil // Prevent further checks on ch2
				}
			}

			// Exit the loop when both channels are closed
			if ch1 == nil && ch2 == nil {
				close(out) // Close the output channel to signal no more messages
				return
			}
		}
	}()
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	out := make(chan string)

	// Start the fanIn function
	fanIn(ch1, ch2, out)

	// Send messages to ch1 and ch2
	go func() {
		ch1 <- "Message 1 from ch1"
		ch1 <- "Message 2 from ch1"
		close(ch1) // Close ch1 after sending messages
	}()

	go func() {
		ch2 <- "Message 1 from ch2"
		ch2 <- "Message 2 from ch2"
		close(ch2) // Close ch2 after sending messages
	}()

	// Receive messages from the output channel
	for msg := range out {
		fmt.Println(msg)
	}

	fmt.Println("All messages processed")
}
