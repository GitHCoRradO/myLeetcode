package main

import (
	"fmt"
	"sort"
)

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		ij := fmt.Sprintf("%d%d", nums[i], nums[j])
		ji := fmt.Sprintf("%d%d", nums[j], nums[i])
		return ij > ji
	})
	if len(nums) != 0 {
		if nums[0] == 0 {
			return "0"
		}
	}
	res := ""
	for _, num := range nums {
		res = res + fmt.Sprintf("%d", num)
	}
	return res
}
