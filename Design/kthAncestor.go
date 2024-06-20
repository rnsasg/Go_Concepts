package main

import "fmt"

func main() {
	//[7, [-1, 0, 0, 1, 1, 2, 2]]
	in := []int{-1, 0, 0, 1, 1, 2, 2}
	//fmt.Println(getKthAncestor(3, 1, in))
	fmt.Println(getKthAncestor(5, 2, in))
	//fmt.Println(getKthAncestor(6, 3, in))
}

func getKthAncestor(node int, k int, in []int) int {
	//fmt.Println(node, k)
	for node > 0 && k > 0 {
		node = node / 2
		k--
		//fmt.Println(node, k)
	}

	if k == 0 {
		return node
	}

	return -1
}
