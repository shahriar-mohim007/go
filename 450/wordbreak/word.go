package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = false
	}
	dp[len(s)] = true
	for i := n - 1; i >= 0; i-- {
		for _, word := range wordDict {
			wordLen := len(word)
			if i+wordLen <= n && s[i:i+wordLen] == word {
				dp[i] = dp[i+wordLen]
			}
			if dp[i] {
				break
			}
		}
	}
	return dp[0]
}

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
}
