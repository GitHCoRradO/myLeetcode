package main

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = 1
	maxlength := 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = maxA_B(dp[i], dp[j]+1)
			}
		}
		maxlength = maxA_B(maxlength, dp[i])
	}
	return maxlength
}

func maxA_B(a, b int) int {
	if a > b {
		return a
	}
	return b
}