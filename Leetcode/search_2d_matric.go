package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	r, c := len(matrix), len(matrix[0])
	left, right := 0, r*c-1

	for left <= right {
		mid := (left + right) / 2
		i, j := mid/c, mid%c

		if matrix[i][j] == target {
			return true
		} else if target < matrix[i][j] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

// func searchMatrix(matrix [][]int, target int) bool {

// 	if len(matrix) == 0 {
// 		return false
// 	}

// 	r := len(matrix)
// 	c := len(matrix[0])
// 	//fmt.Println("Row: ", r, " Column: ", c)

// 	if target < matrix[0][0] || target > matrix[r-1][c-1] {
// 		return false
// 	}

// 	rlow := 0
// 	rhigh := r - 1
// 	row := 0
// 	//fmt.Println("RLow: ", rlow, " RHigh: ", rhigh)

// 	for rlow <= rhigh {
// 		rmid := rlow + (rhigh-rlow)/2
// 		//fmt.Println("RLow: ", rlow, " RHigh: ", rhigh, " RMid: ", rmid)
// 		if target == matrix[rmid][0] || target == matrix[rmid][c-1] {
// 			return true
// 		}

// 		if target > matrix[rmid][0] && target < matrix[rmid][c-1] {
// 			row = rmid
// 			break
// 		} else if target < matrix[rmid][0] {
// 			//fmt.Println("Changing rhigh")
// 			rhigh = rmid - 1
// 		} else {
// 			//fmt.Println("Changing rlow")
// 			rlow = rmid + 1
// 		}
// 	}

// 	low := 0
// 	high := c - 1
// 	fmt.Println("Row :", row)
// 	for low <= high {
// 		mid := low + (high-low)/2
// 		//fmt.Println("Low: ", low, " High: ", high, " Mid: ", mid)
// 		if target == matrix[row][mid] {
// 			return true
// 		} else if target > matrix[row][mid] {
// 			low = mid + 1
// 		} else {
// 			high = mid - 1
// 		}
// 	}

// 	return false
// }

func main() {
	fmt.Println(searchMatrix([][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}, 3))
}
