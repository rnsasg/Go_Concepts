package main

import (
	"fmt"
	"sort"
)

type TwoDArray [][]int

func (a TwoDArray) Len() int           { return len(a) }
func (a TwoDArray) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a TwoDArray) Less(i, j int) bool { return a[i][0] < a[j][0] }

func merge(intervals [][]int) [][]int {

	if len(intervals) == 1 {
		return intervals
	}
	var arr TwoDArray = intervals
	sort.Sort(arr)

	intervals = arr
	fmt.Println(intervals)
	var res [][]int
	start := intervals[0][0]
	end := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if end >= intervals[i][0] && end >= intervals[i][1] {
			continue
		}
		if end >= intervals[i][0] {
			end = intervals[i][1]
			continue
		} else {
			res = append(res, []int{start, end})
			start = intervals[i][0]
			end = intervals[i][1]
		}
	}
	res = append(res, []int{start, end})
	return res
}

func main() {
	//var arr TwoDArray = [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	//var arr TwoDArray = [][]int{{1, 4}, {0, 4}}
	var arr TwoDArray = [][]int{{1, 4}, {2, 3}}
	//sort.Sort(arr)
	// fmt.Println(arr)
	fmt.Println(merge(arr))
}
