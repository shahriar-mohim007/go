package main

import (
	"fmt"
	"sync"
)

var once sync.Once
var instance *MySingleton

// MySingleton is a simple example of a singleton class
type MySingleton struct {
	Name string
}

// GetInstance returns the singleton instance
func GetInstance() *MySingleton {
	once.Do(func() {
		instance = &MySingleton{Name: "Singleton Instance"}
		fmt.Println("instance:", instance)
	})
	return instance
}

func main() {
	// Fetch the singleton instance
	instance1 := GetInstance()
	instance2 := GetInstance()

	// Both instances should point to the same memory address
	fmt.Println(instance1) // &{Singleton Instance}
	fmt.Println(instance2) // &{Singleton Instance}

	// Check if both are the same instance
	fmt.Println(instance1 == instance2) // true
}
