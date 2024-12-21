package main

import "fmt"

func longestPalindrome(s string) string {
	var res string
	reslen := 0

	for i := 0; i < len(s); i++ {
		// Odd-length palindrome (centered at i)
		l, r := i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > reslen {
				reslen = r - l + 1
				res = s[l : r+1]
			}
			l--
			r++
		}

		// Even-length palindrome (centered between i and i+1)
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > reslen {
				reslen = r - l + 1
				res = s[l : r+1]
			}
			l--
			r++
		}
	}
	return res
}

func main() {
	fmt.Println(longestPalindrome("aab"))
}
