package main

import "fmt"

func lengthOfLongestSubstring2(s string) int {
	left, curr, ans := 0, 0, 0
	visited := make(map[byte]bool)

	for right := 0; right < len(s); right++ {
		for visited[s[right]] {
			delete(visited, s[left])
			left++
		}
		visited[s[right]] = true
		curr = right - left + 1
		ans = max(ans, curr)
	}
	return ans
}
func lengthOfLongestSubstring(s string) int {
	left, ans := 0, 0
	lasindex := [256]int{}
	for i := range lasindex {
		lasindex[i] = -1
	}
	for right := 0; right < len(s); right++ {
		if lasindex[s[right]] >= left {
			left = lasindex[s[right]] + 1
		}
		lasindex[s[right]] = right
		ans = max(ans, right-left+1)
	}
	return ans
}
func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
