package main

import "fmt"

// Ref: https://clavinjune.dev/en/blogs/max-heap-in-go/
type MaxHeap struct {
	heap     []int
	capacity int
	size     int
	root     int
}

func (MaxHeap) getParent(idx int) int {
	return idx / 2
}

func (MaxHeap) getLeft(idx int) int {
	return idx * 2
}

func (MaxHeap) getRight(idx int) int {
	return idx*2 + 1
}

func (m *MaxHeap) swap(idx1, idx2 int) {
	m.heap[idx1], m.heap[idx2] = m.heap[idx2], m.heap[idx1]
}

func (m MaxHeap) String() string {
	return fmt.Sprintf(`size:     %v
capacity: %v
heap:     %v`,
		m.size, m.capacity-1, m.heap[1:])
}

func (m *MaxHeap) rebalance(idx int) {
	if idx >= m.size {
		return
	}

	left := m.getLeft(idx)
	right := m.getRight(idx)

	if left <= m.size {
		if m.heap[idx] < m.heap[left] {
			m.swap(idx, left)
			m.rebalance(left)
		}
	}

	if right <= m.size {
		if m.heap[idx] < m.heap[right] {
			m.swap(idx, right)
			m.rebalance(right)
		}
	}
}

func (m *MaxHeap) Pop() (int, error) {
	if m.size <= 0 {
		return (1 << 31) - 1, fmt.Errorf("max-heap: heap is empty")
	}

	// fetch the root
	max := m.heap[m.root]
	// make the route is the last element in the heap
	m.heap[m.root] = m.heap[m.size]
	// make it zero value
	m.heap[m.size] = 0
	// decrease the size so the zero value won't be counted
	m.size--

	// rebalance the heap
	m.rebalance(m.root)

	// return the greatest
	return max, nil
}

func (m MaxHeap) Peek() (int, error) {
	if m.size <= 0 {
		return (1 << 31) - 1, fmt.Errorf("max-heap: heap is empty")
	}

	return m.heap[m.root], nil
}

func (m *MaxHeap) Push(element int) {
	// exceed the limit
	if m.size >= m.capacity-1 {
		return
	}

	m.size++
	idx := m.size

	m.heap[idx] = element

	parent := m.getParent(idx)

	for m.heap[idx] > m.heap[parent] {
		m.swap(idx, parent)

		idx = parent
		parent = m.getParent(idx)
	}
}

func NewMaxHeap(capacity int) *MaxHeap {
	// because we used the 0 index
	// we need to increase the capacity defined by user
	c := capacity + 1
	h := make([]int, c, c)

	// just to mark that it is the minimum one,
	// you can ignore this
	h[0] = (1 << 31) - 1

	return &MaxHeap{
		root:     1, // always 1, you can omit this attribute
		size:     0,
		capacity: c,
		heap:     h,
	}
}

func main() {
	h := NewMaxHeap(15)
	h.Push(1)
	h.Push(4)
	h.Push(2)
	h.Push(5)
	h.Push(3)
	h.Push(10)
	fmt.Println(h)
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	h.Push(11)

	fmt.Println(h)
}
