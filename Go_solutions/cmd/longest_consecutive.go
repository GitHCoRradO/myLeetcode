package main

func longestConsecutive(nums []int) int {
	numsSet := make(map[int]bool)
	for _, i := range nums {
		numsSet[i] = true
	}
	longestConsecutive := 0
	for i, _ := range numsSet {
		if !numsSet[i-1] {
			currentNum := i
			currentLength := 1
			for numsSet[currentNum+1] {
				currentNum += 1
				currentLength += 1
			}
			if longestConsecutive < currentLength {
				longestConsecutive = currentLength
			}
		}
	}
	return longestConsecutive
}
