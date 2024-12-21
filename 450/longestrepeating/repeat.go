package main

import "fmt"

func characterReplacement(s string, k int) int {
	count := make(map[byte]int)
	maxCount, start, maxLength := 0, 0, 0
	for end := 0; end < len(s); end++ {
		count[s[end]]++
		if count[s[end]] > maxCount {
			maxCount = count[s[end]]
		}
		if (end-start+1)-maxCount > k {
			count[s[start]]--
			start++
		}
		maxLength = max(maxLength, end-start+1)
	}
	return maxLength
}

func main() {
	fmt.Println(characterReplacement("AABABBA", 1))
}
