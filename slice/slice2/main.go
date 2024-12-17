package main

import "fmt"

func main() {
	// Create a slice with an initial length of 3 and capacity of 5
	slice := make([]int, 3)
	//fmt.Println(slice[0])panic: runtime error: index out of range [0] with length 0
	fmt.Println("Initial slice:", slice)
	fmt.Println("Length:", len(slice), "Capacity:", cap(slice))

	// Append more elements, which will trigger a resize once capacity is exceeded
	slice = append(slice, 10, 20, 30, 40, 50, 60)
	fmt.Println("After appending:", slice)
	fmt.Println("Length:", len(slice), "Capacity:", cap(slice))

	//Initial slice: [0 0 0]
	//Length: 3 Capacity: 5
	//After appending: [0 0 0 10 20 30]
	//Length: 6 Capacity: 10
}
