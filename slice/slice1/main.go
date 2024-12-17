package main

import (
	"fmt"
	"slices"
)

func main() {

	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	shared := s
	shared[0] = "mohim"
	fmt.Println("cpy:", shared, "main", s)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	twoD := make([][]int, 3, 4) // Length 3, capacity 4
	fmt.Println("len:", len(twoD), "cap:", cap(twoD))

	for i := 0; i < 4; i++ { // Looping to 4
		innerLen := i + 1
		if i >= len(twoD) { // Extend the slice if index is out of range
			twoD = append(twoD, make([]int, innerLen))
		} else {
			twoD[i] = make([]int, innerLen)
		}

		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	var x []int
	x = append(x, 10)
	fmt.Println(x)
	x = append(x, 5, 6, 7)
	fmt.Println(x)
	y := []int{20, 30, 40}
	x = append(x, y...)
	fmt.Println(x)
	//y... is a variadic argument that expands the slice y into its individual elements.
	//	This means you're appending all elements of y to x in one operation.

	//slice := make([]T, length, capacity)
	//T is the type of the elements (e.g., int, string, etc.).
	//length is the initial number of elements in the slice.
	//capacity is the total number of elements the slice can hold without reallocation.

	//Go is a call by value
	//language. Every time you pass a parameter to a function, Go makes a copy
	//of the value that’s passed in. Passing a slice to the append function actually
	//passes a copy of the slice to the function. The function adds the values to
	//the copy of the slice and returns the copy. You then assign the returned slice
	//back to the variable in the calling function.

}
