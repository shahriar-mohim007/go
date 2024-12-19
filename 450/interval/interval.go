package main

import (
	"fmt"
	"sort"
)

func insert(intervals [][]int, newInterval []int) [][]int {
	// Append the new interval
	intervals = append(intervals, newInterval)

	// Sort intervals by their start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	mergedIntervals := make([][]int, 0)

	for _, interval := range intervals {
		// If mergedIntervals is empty or there is no overlap
		if len(mergedIntervals) == 0 || mergedIntervals[len(mergedIntervals)-1][1] < interval[0] {
			mergedIntervals = append(mergedIntervals, interval)
		} else {
			// There is overlap, merge the intervals
			mergedIntervals[len(mergedIntervals)-1][1] = max(mergedIntervals[len(mergedIntervals)-1][1], interval[1])
		}
	}

	return mergedIntervals
}

func main() {
	intervals := [][]int{
		{1, 3},
		{6, 9},
	}
	newInterval := []int{2, 5}

	result := insert(intervals, newInterval)
	for _, interval := range result {
		fmt.Println(interval)
	}
}
