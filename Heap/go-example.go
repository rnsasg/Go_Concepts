package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (m MinHeap) Len() int {
	return len(m)
}

func (m MinHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m MinHeap) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m *MinHeap) Pop() any {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}

func (m *MinHeap) Push(x any) {
	*m = append(*m, x.(int))
}

func main() {
	nums := &MinHeap{3, 1, 4, 5, 1, 1, 2, 6}
	heap.Init(nums)
	fmt.Println(nums)

	min := heap.Pop(nums)
	fmt.Println("min: ", min, " heap: ", *nums)

	heap.Push(nums, 5)
	fmt.Println("heap: ", *nums)

	sortedNums := []int{}
	for nums.Len() > 0 {
		sortedNums = append(sortedNums, heap.Pop(nums).(int))
	}
	fmt.Println("sorted values: ", sortedNums)
}
