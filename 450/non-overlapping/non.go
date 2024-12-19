package main

import (
	"fmt"
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] == intervals[j][1] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
	})
	second := intervals[0][1]
	count := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= second {
			second = intervals[i][1]
		} else {
			count++
		}
	}
	return count
}

func main() {
	intervals := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{1, 3},
	}

	fmt.Println("Input intervals:", intervals)
	overlapCount := eraseOverlapIntervals(intervals)
	fmt.Println("Number of intervals to remove to avoid overlap:", overlapCount)
}
