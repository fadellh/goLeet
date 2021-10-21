package main

import (
	"math"
)

func main() {
	// fmt.Println(SearchInsert([]int{1, 3, 5, 7}, 4))
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	// test := "halo"
	// // QuickSort(arr, 0, len(arr)-1)
	// fmt.Println(&test)
	// fmt.Println(&arr[1])
	Rotate(arr, 3)
}

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		middle := (left + right) / 2
		middle = int(math.Floor(float64(middle)))

		if target > nums[middle] {
			left = middle
			right--
		}

		if target < nums[middle] {
			right = middle
			left++
		}

		if target == nums[middle] {
			return nums[middle]
		}
	}
	return -1
}
