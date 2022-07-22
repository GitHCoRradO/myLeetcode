package main

import "math"

func findPeakElement(nums []int) int {
	n := len(nums)
	get := func(i int) int {
		if i == -1 || i == n {
			return math.MinInt64
		}
		return nums[i]
	}
	low := 0
	high := n - 1
	for {
		mid := (low + high) / 2
		if (get(mid) > get(mid-1)) && (get(mid) < get(mid+1)) {
			return mid
		}
		if get(mid) < get(mid + 1) {
			low = mid + 1
		} else {
			high = mid -1
		}
	}
}
