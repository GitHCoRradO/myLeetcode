package main

import "math"

/*
f(i) = 1 + min (f(i - j*j))  ( j : [1, i^1/2])
 */
func numSquares(n int) int {
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		min := math.MaxInt32
		for j := 1; j * j <= i; j++ {
			min = minAB(min, f[i - j*j])
		}
		f[i] = min + 1
	}
	return f[n]
}

func minAB(a, b int) int {
	if a < b {
		return a
	}
	return b
}