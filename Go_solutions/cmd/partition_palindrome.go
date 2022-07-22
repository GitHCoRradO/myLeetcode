package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {
	input := os.Args[1]
	sli := partition(input)
	fmt.Printf("result set: %v\n", sli)
	for _, s := range sli {
		fmt.Println(s)
	}
}

var resultSet [][]string

func partition(s string) [][]string {
	resultSet = make([][]string, 0)
	for i := 0; i < len(s); i++ {
		partitionBacktrack(0, i, s, []string{})
	}
	return resultSet
}

func partitionBacktrack(startIndex, endIndex int,  s string, ans []string) {
	fmt.Printf("start %d end %d\n", startIndex, endIndex)
	fmt.Printf("#1. ans: %v\n", ans)
	if endIndex == len(s) {

		return
	}
	subStr := s[startIndex:endIndex+1]
	if isPalindrome(subStr) {
		ans = append(ans, subStr)
		fmt.Printf("#2. ans: %v\n", ans)
		for i := endIndex+1; i <= len(s); i++ {
			subAns := make([]string, 0)
			partitionBacktrack(endIndex+1, i, s, subAns)
		}
	} else {
		fmt.Printf("#3. ans: %v\n", ans)
		partitionBacktrack(startIndex, endIndex+1, s, ans)
		resultSet = append(resultSet, ans)
	}
}

func isPalindrome(s string) bool {
	newSlice := make([]rune, 0)
	for _, c := range s {
		if unicode.IsDigit(c) || unicode.IsLetter(c) {
			newSlice = append(newSlice, unicode.ToLower(c))
		}
	}
	low := 0
	high := len(newSlice) - 1
	for ;low <= high; {
		if newSlice[low] == newSlice[high] {
			low += 1
			high -= 1
			continue
		} else {
			return false
		}
	}
	return true
}



