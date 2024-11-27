package main

import "fmt"

func main() {
	// Declare a map using make
	myMap := make(map[string]int)
	v, ok := myMap["hello"]
	fmt.Println(v, ok)
	// Add key-value pairs
	myMap["Alice"] = 25
	myMap["Bob"] = 30

	// Print the map
	fmt.Println("Map:", myMap) // Map: map[Alice:25 Bob:30]

	totalWins := map[string]int{}
	v, ok = totalWins["Alice"]
	fmt.Println(v, ok)
	totalWins["Bob"] = 10

	fmt.Println("myMap:", myMap)         // myMap: map[Alice:5]
	fmt.Println("totalWins:", totalWins) // totalWins: map[Bob:10]
	m := map[string]int{
		"hello": 5,
		"world": 10,
	}
	delete(m, "hello")
	fmt.Println(m)
}
