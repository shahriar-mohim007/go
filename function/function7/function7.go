// package main
//
// import "fmt"
//
//	func makeMult(base int) func(int) int {
//		return func(factor int) int {
//			return base * factor
//		}
//	}
//
//	func main() {
//		twoBase := makeMult(2)
//		threeBase := makeMult(3)
//		for i := 0; i < 3; i++ {
//			fmt.Println(twoBase(i), threeBase(i))
//		}
//	}
package main

import "fmt"

func modifySlice(s []int) {
	s[0] = 100 // Modify the first element
}

func main() {
	slice := []int{1, 2, 3}
	modifySlice(slice)
	fmt.Println(slice) // Output: [100 2 3]
}
