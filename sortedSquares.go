package main

import "math"

func SortedSquares(nums []int) []int {

	output := []int{}
	var temp int

	for _, item := range nums {
		temp = item * item
		output = append(output, temp)
	}

	QuickSort(output, 0, len(nums)-1)

	return output
}

func SortSquresLeet(nums []int) []int {

	left := 0
	n := len(nums)
	right := n - 1
	result := nums
	square := 0

	for i := n - 1; i >= 0; i-- {
		if math.Abs(float64(nums[left])) < math.Abs(float64(nums[right])) {
			square = nums[right]
			right--
		} else {
			square = nums[left]
			left++
		}
		result[i] = square * square
	}

	return result

}

func QuickSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	//Partitiion
	boundary := partition(arr, start, end)
	//Sort Left
	QuickSort(arr, start, boundary-1)
	//Sort Right
	QuickSort(arr, boundary+1, end)

	return
}

func partition(arr []int, start, end int) int {
	pivot := arr[end]
	boundary := start - 1

	for i := start; i <= end; i++ {
		if arr[i] <= pivot {
			boundary++
			swap(arr, i, boundary)
		}
	}

	return boundary
}

func swap(arr []int, idx1, idx2 int) {
	temp := arr[idx1]
	arr[idx1] = arr[idx2]
	arr[idx2] = temp
}
