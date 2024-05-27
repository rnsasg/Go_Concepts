package main

import "fmt"

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}

func search(nums []int, target int) int {
	pivot := findPivot(nums, 0, len(nums)-1)
	fmt.Println("Pivot element", pivot+1)
	if pivot == -1 {
		return binarySearch(nums, 0, len(nums)-1, target)
	}

	if res := binarySearch(nums, 0, pivot-1, target); res != -1 {
		return res
	}
	return binarySearch(nums, pivot, len(nums)-1, target)
}

// func searchPivot(nums []int, low, high int) int {

// 	for low <= high {
// 		mid := (low + high) >> 1
// 		if nums[mid] < nums[mid-1] {
// 			return mid
// 		} else if nums[mid] >= nums[mid-1] {
// 			low = mid + 1
// 		} else if nums[mid] < nums[mid+1] {
// 			high = mid - 1
// 		}
// 	}

//		return -1
//	}
func findPivot(nums []int, low, high int) int {
	for low < high {
		mid := (low + high) >> 1

		if nums[mid] > nums[mid+1] {
			return mid
		} else if nums[mid] < nums[mid+1] {
			return mid - 1
		} else if nums[low] >= nums[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func binarySearch(nums []int, low, high, target int) int {
	for low <= high {
		mid := (low + high) >> 1
		if nums[mid] == target {
			return mid
		} else if target < nums[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
