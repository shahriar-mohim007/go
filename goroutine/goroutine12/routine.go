package main

import (
	"errors"
	"fmt"
)

type PressureGauge struct {
	ch chan struct{}
}

// New creates a new PressureGauge with a specified limit.
func New(limit int) *PressureGauge {
	ch := make(chan struct{}, limit)
	// Fill the channel with 'struct{}' values to represent available capacity.
	for i := 0; i < limit; i++ {
		ch <- struct{}{}
	}
	return &PressureGauge{
		ch: ch,
	}
}

// Process allows a function to run if there is available capacity in the PressureGauge.
func (pg *PressureGauge) Process(f func()) error {
	select {
	case <-pg.ch:
		// There is capacity, run the function.
		f()
		// After processing, signal that one spot is free.
		pg.ch <- struct{}{}
		return nil
	default:
		// No capacity available, return an error.
		return errors.New("no more capacity")
	}
}

func main() {
	// Create a new PressureGauge with a limit of 2.
	pg := New(2)

	// Define a sample function to process.
	processFunction := func() {
		fmt.Println("Processing...")
	}

	// Try to process a few times.
	for i := 0; i < 5; i++ {
		err := pg.Process(processFunction)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}
