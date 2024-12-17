package main

import "fmt"

func main() {
	var s string = "Hello there"
	var b byte = s[6]
	fmt.Println(b)

	var s2 string = s[4:7]
	var s3 string = s[:5]
	var s4 string = s[6:]
	fmt.Println(s2, s3, s4)

	var a rune = 'x'
	var s1 string = string(a)
	var b1 byte = 'y'
	var s5 string = string(b1)
	fmt.Println(s1, s5, b1)

	var s6 string = "Hello"
	var bs []byte = []byte(s6)
	var rs []rune = []rune(s6)
	fmt.Println(bs)
	fmt.Println(rs)
}

//Go uses a sequence of bytes to represent a
//string.
