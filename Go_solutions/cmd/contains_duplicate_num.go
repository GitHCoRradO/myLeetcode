package main

func containsDuplicate(nums []int) bool {
	intMap := make(map[int]bool)
	for _, value := range nums {
		if _, exists := intMap[value]; exists {
			return true
		} else {
			intMap[value] = true
		}
	}
	return false
}
