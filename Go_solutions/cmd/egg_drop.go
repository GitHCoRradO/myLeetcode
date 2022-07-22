package main

import "fmt"

func main() {
	fmt.Println(superEggDrop(2, 3))
}

func superEggDrop(k int, n int) int {
	dp := make([][]int, k + 1)
	for i := range dp {
		dp[i] = make([]int, n + 1)
	}
	m := 0
	for dp[k][m] < n {
		m++
		for i := 1; i <= k; i++ {
			dp[i][m] = 1 + dp[i][m-1] + dp[i-1][m-1]
		}
	}
	return m
}