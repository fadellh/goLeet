package main

func SearchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	middle := 0

	for left <= right {
		middle = (left + right) / 2
		// fmt.Println(middle)

		if target == nums[middle] {
			return middle
		}
		if target > nums[middle] {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	if target > nums[middle] {
		return middle + 1
	} else {
		return middle
	}

}
