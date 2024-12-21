package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	Key   int
	Value int
}

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, n := range nums {
		freq[n]++
	}
	pairs := make([]Pair, 0, len(freq))
	for key, v := range freq {
		pairs = append(pairs, Pair{key, v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value > pairs[j].Value
	})
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = pairs[i].Key
	}
	return res
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(topKFrequent(nums, 2))
}
