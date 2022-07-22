package main

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	nums1nums2Sum := map[int]int{}
	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			nums1nums2Sum[num1+num2] += 1
		}
	}
	cnt := 0
	for _, num3 := range nums3 {
		for _, num4 := range nums4 {
			num3num4Sum := 0 - (num3 + num4)
			value, exists := nums1nums2Sum[num3num4Sum]
			if exists {
				cnt += value
			}
		}
	}
	return cnt
}
