package main

import (
	"fmt"
	"unicode"
)

func main() {
	str := "A man, a plan, a canal: Panama"
	fmt.Println(str)
	fmt.Println(IsPalindrome(str))
}

func IsPalindrome(s string) bool {
	newSlice := make([]rune, 0)
	for _, c := range s {
		if unicode.IsDigit(c) || unicode.IsLetter(c) {
			newSlice = append(newSlice, unicode.ToLower(c))
		}
	}
	fmt.Printf("new string:\n%s\n", string(newSlice))
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



