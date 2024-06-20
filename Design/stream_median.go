package main

import "fmt"

type MedianFinder struct {
	arr    []int
	median float64
}

func Constructor() MedianFinder {
	arr := make([]int, 0)
	return MedianFinder{
		arr:    arr,
		median: 0.0,
	}
}

func (this *MedianFinder) AddNum(num int) {

	if len(this.arr) == 0 || this.arr[len(this.arr)-1] <= num {
		this.arr = append(this.arr, num)
	} else if this.arr[0] >= num {
		this.arr = append([]int{num}, this.arr...)
	} else {
		i := 0
		for ; i < len(this.arr) && this.arr[i] < num; i++ {
		}
		// temp := this.arr[i:]
		temp := make([]int, len(this.arr[i:]))
		copy(temp, this.arr[i:])
		fmt.Println(temp)
		this.arr = append(this.arr[:i], num)
		fmt.Println(this.arr)
		this.arr = append(this.arr, temp...)
		fmt.Println(this.arr)
	}

	if len(this.arr)%2 == 0 {
		mid1 := float64(this.arr[len(this.arr)/2])
		mid2 := float64(this.arr[(len(this.arr)/2)-1])
		this.median = float64((mid1 + mid2) / 2)
		//fmt.Println(this.arr)
	} else {
		this.median = float64(this.arr[len(this.arr)/2])
		//fmt.Println(this.arr)
	}
}

func (this *MedianFinder) FindMedian() float64 {

	if len(this.arr) == 0 {
		return -1
	}
	return this.median
}

func main() {
	obj := Constructor()
	obj.AddNum(6)
	obj.FindMedian()
	obj.AddNum(10)
	obj.FindMedian()
	obj.AddNum(2)
	obj.FindMedian()
	obj.AddNum(6)
	obj.FindMedian()
	obj.AddNum(5)
	obj.FindMedian()
}

// ["MedianFinder","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian"]
// [[],[6],[],[10],[],[2],[],[6],[],[5],[]]
