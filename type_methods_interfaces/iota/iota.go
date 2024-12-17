package main

import "fmt"

const (
	A = iota // 0
	B        // 1
	C        // 2
)

const (
	KB = 1 << (10 * iota) // 1 << (10*0) = 1
	MB                    // 1 << (10*1) = 1024
	GB                    // 1 << (10*2) = 1048576
)

const (
	First  = iota // 0
	_             // 1 (skipped)
	Second        // 2
)

func main() {
	fmt.Println(A, B, C) // Output: 0 1 2
	fmt.Println(KB, MB, GB)
	fmt.Println(First, Second)
}
