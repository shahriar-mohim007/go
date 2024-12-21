package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)
	for _, str := range strs {
		sortedStr := sortStrings(str)
		anagrams[sortedStr] = append(anagrams[sortedStr], str)
	}
	var result [][]string
	for _, group := range anagrams {
		result = append(result, group)
	}
	return result
}

func sortStrings(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tsr"}))
}
