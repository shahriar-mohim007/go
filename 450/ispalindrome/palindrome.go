package main

import "unicode"

func isPalindrome(s string) bool {
	var filtered []rune
	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			filtered = append(filtered, unicode.ToLower(char))
		}
	}
	n := len(filtered)
	for i := 0; i < n/2; i++ {
		if filtered[i] != filtered[n-i-1] {
			return false
		}
	}
	return true
}
