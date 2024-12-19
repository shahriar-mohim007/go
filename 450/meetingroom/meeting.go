package main

import (
	"fmt"
	"sort"
)

func canAttendMeetings(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			return false
		}
	}
	return true
}

func main() {
	intervals := [][]int{
		{0, 30},
		{5, 10},
		{15, 20},
	}

	canAttend := canAttendMeetings(intervals)
	fmt.Println("Can attend all meetings:", canAttend)
}
