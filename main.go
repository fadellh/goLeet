package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(SearchInsert([]int{1, 3, 5, 7}, 4))
	// fmt.Println(strconv.Itoa(97))
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
