package main

import "strings"

func longestSubstring(s string, k int) (ans int) {
	if s == "" {
		return 0
	}
	letterTimes := make([]int, 26)
	for _, c := range s {
		letterTimes[c-'a']++
	}
	split := ""
	for _, c := range s {
		if letterTimes[c-'a'] > 0 && letterTimes[c-'a'] < k {
			split = string(c)
			break
		}
	}
	if split == "" {
		return len(s)
	}
	for _, subStr := range strings.Split(s, split) {
		ans = maxAB(ans, longestSubstring(subStr, k))
	}
	return
}

func maxAB(a, b int) int {
	if a > b {
		return a
	}
	return b
}