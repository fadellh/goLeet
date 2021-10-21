package main

func Rotate(nums []int, k int) {
	result := make([]int, len(nums))
	copy(result, nums)
	index := 0
	for i := 0; i < len(result); i++ {
		index = i + k
		if index > len(result)-1 {
			index = (index % len(result))
			nums[index] = result[i]
		}
		nums[index] = result[i]
	}
}
