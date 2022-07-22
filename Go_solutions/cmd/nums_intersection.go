package main

func intersect(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	searchedIndices := make(map[int]bool)
	for _, elem1 := range nums1 {
		for i, elem2 := range nums2 {
			if elem1 == elem2 {
				if _, exists := searchedIndices[i]; exists {
					continue
				} else {
					res = append(res, elem1)
					searchedIndices[i] = true
					break
				}
			}
		}
	}
	return res
}