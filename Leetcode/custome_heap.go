package main

import (
	"container/heap"
	"fmt"
)

type Pair struct {
	first  int
	second int
}

type MinHeap []Pair

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return (h[i].first + h[i].second) < (h[j].first + h[j].second) }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Pair)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}
func (p Pair) String() string {
	return fmt.Sprintf("[ %d, %d ]", p.first, p.second)
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	harr := &MinHeap{}
	// Create pairs
	heap.Init(harr)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			heap.Push(harr, Pair{first: nums1[i], second: nums2[j]})
		}
	}
	// var ele Pair
	var res [][]int
	//fmt.Println(*harr)

	for i := k; i > 0; i-- {
		ele := heap.Pop(harr).(Pair)
		res = append(res, []int{ele.first, ele.second})
	}

	return res
}

func main() {
	//fmt.Println(kSmallestPairs([]int{1, 7, 11}, []int{2, 4, 6}, 3))
	fmt.Println(kSmallestPairs([]int{1, 1, 2}, []int{1, 2, 3}, 3))
}
