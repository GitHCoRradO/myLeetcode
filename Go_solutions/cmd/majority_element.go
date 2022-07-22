package main

func majorityElement(nums []int) int {
	elemCnt := make(map[int]int)
	for _, value := range nums {
		if _, exists := elemCnt[value]; !exists {
			elemCnt[value] = 1
		} else {
			elemCnt[value] += 1
		}
	}
	for elem, cnt := range elemCnt {
		if cnt > len(nums) / 2  {
			return elem
		}
	}
	return 0
}