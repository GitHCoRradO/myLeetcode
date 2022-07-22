package main

import "sort"

/*
Given an integer array nums and an integer k, return the kth largest element in the array.

Note that it is the kth largest element in the sorted order, not the kth distinct element.

 */
func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	index := len(nums) - k
	if index < 0 {
		return nums[0]
	} else {
		return nums[index]
	}
}
