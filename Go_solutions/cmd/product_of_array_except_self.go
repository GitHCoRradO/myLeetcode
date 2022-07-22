package main

func productExceptSelf(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	lProducts := make([]int, len(nums))
	lProducts[0] = 1
	rProducts := make([]int, len(nums))
	rProducts[len(nums)-1] = 1
	for lIndex := 1; lIndex < len(lProducts); lIndex++ {
		lProducts[lIndex] = lProducts[lIndex - 1]*nums[lIndex-1]
	}
	for rIndex := len(rProducts) - 2; rIndex > 0 ; rIndex++ {
		rProducts [rIndex] = rProducts[rIndex+1]*nums[rIndex+1]
	}
	res := make([]int, len(nums))
	for i := 0; i < len(res); i++ {
		res[i] = lProducts[i]*rProducts[i]
	}
	return res
}
