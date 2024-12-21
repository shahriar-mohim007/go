package main

import "fmt"

func Palindrome(s string, l, r int) int {
	res := 0
	for l >= 0 && r < len(s) && s[l] == s[r] {
		res++
		l--
		r++
	}
	return res
}

func countSubstrings(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		res += Palindrome(s, i, i)
		res += Palindrome(s, i, i+1)
	}
	return res
}

func main() {
	fmt.Println(countSubstrings("abc"))
}
