package main

import (
	"fmt"
	"sync"
)

var once sync.Once
var resource string

// This function will only be executed once.
func initializeResource() {
	resource = "Initialized Resource"
	fmt.Println("Resource initialized.")
}

func main() {
	// Call the initialize function multiple times.
	once.Do(initializeResource) // This will initialize the resource
	once.Do(initializeResource) // This will be ignored
	once.Do(initializeResource) // This will be ignored

	// Access the initialized resource
	fmt.Println("Resource:", resource)
}

//In Go, sync.Once is a type in the sync package that
//ensures a function is only executed once, even if it is called multiple times across different goroutines.
//It is commonly used for lazy initialization or to ensure that a specific action is done only once, regardless of
//how many times the function or method is called.
//Key Features:
//
//Thread-safe: It ensures that the function is only executed once, even when multiple goroutines are involved.
//Idempotent: The function will be executed only the first time, and subsequent calls will not invoke it again.
//Useful for initialization: It is often used to initialize global variables, configuration setups, or resources that should only be initialized once.
//
//How It Works:
//
//sync.Once has a method called Do, which takes a function as its argument.
//The first time Do is called, it executes the function provided. On subsequent calls,
//it ensures that the function is not called again.
//If the function passed to Do panics or has an error, sync.Once will not attempt to call the function again.
