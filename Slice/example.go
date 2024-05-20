package main

import "fmt"

func main() {
	arr := []int{45, 32, 12, 78, 13, 36, 56}
	x := 1
	temp := arr[x+1]
	y := 4
	t := arr[y+1:]
	arr = append(arr[:x+1], arr[x+2:y+1]...)
	arr = append(arr, temp)
	arr = append(arr, t...)
	fmt.Println(arr)
}
