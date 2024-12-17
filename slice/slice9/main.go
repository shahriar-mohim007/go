package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	y := make([]int, 2)
	copy(y, x[2:])
	fmt.Println(x)
	fmt.Println(y)

	num := copy(x[:3], x[1:])
	fmt.Println(x, num)

	d := [4]int{5, 6, 7, 8}
	Y := make([]int, 2)
	copy(Y, d[:])
	fmt.Println(Y)
	copy(d[:], x)
	fmt.Println(d)

	//d := [4]int{5, 6, 7, 8}
	//y := make([]int, 2)
	//copy(y, d[:])
	//fmt.Println(y)
	//copy(d[:], x)
	//fmt.Println(d)

}

//[1 2 3 4]
//[3 4]
//[2 3 4 4] 3
//[5 6]
//[2 3 4 4]
