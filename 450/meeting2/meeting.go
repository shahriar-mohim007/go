package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minMeetingRooms(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	minheap := IntHeap{}
	heap.Init(&minheap)
	for _, interval := range intervals {
		if minheap.Len() > 0 && interval[0] < minheap[0] {
			heap.Pop(&minheap)
		}
		heap.Push(&minheap, interval[1])
	}
	return minheap.Len()
}

func main() {
	intervals := [][]int{
		{0, 40},
		{5, 10},
		{15, 20},
	}
	fmt.Println("Input intervals:", intervals)
	rooms := minMeetingRooms(intervals)
	fmt.Println("Minimum number of meeting rooms required:", rooms)
}
