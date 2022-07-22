package main

import "math"

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	dp[0] = 0
	for j := 1; j <= amount; j++ {
		dp[j] = math.MaxInt64
		for i := 0; i < len(coins); i++ {
			if j >= coins[i] && dp[j-coins[i]] != math.MaxInt64 {
				dp[j] = minab(dp[j], dp[j-coins[i]]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt64 {
		return -1
	}
	return dp[amount]
}

func minab(a, b int) int {
	if a < b {
		return a
	}
	return b
}
