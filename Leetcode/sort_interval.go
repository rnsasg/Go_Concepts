package main

import (
	"fmt"
	"sort"
)

type TwoDArray [][]int

func (a TwoDArray) Len() int           { return len(a) }
func (a TwoDArray) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a TwoDArray) Less(i, j int) bool { return a[i][0] < a[j][0] }

func main() {
	//var arr TwoDArray = [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	//var arr TwoDArray = [][]int{{1, 4}, {0, 4}}
	var arr TwoDArray = [][]int{{1, 4}, {2, 3}}
	sort.Sort(arr)
	fmt.Println(arr)
}
