package main

import "fmt"

func main() {
	nums := []int{1, 0, 1}
	moveZeroes(nums)
	fmt.Printf("%v\n", nums)
}

func moveZeroes(nums []int)  {
	if len(nums) == 0 {
		return
	}
	zeroPointer, nonZeroPointer := 0, 0
	for nonZeroPointer < len(nums) {
		for nonZeroPointer < len(nums) {
			if nums[nonZeroPointer] == 0 {
				nonZeroPointer++
			} else {
				break
			}
		}
		for zeroPointer < len(nums) {
			if nums[zeroPointer] != 0 {
				zeroPointer++
			} else {
				break
			}
		}
		if nonZeroPointer == len(nums) {
			break
		}

		if zeroPointer < nonZeroPointer {
			nums[zeroPointer] = nums[nonZeroPointer]
			nums[nonZeroPointer] = 0
		}
		//zeroPointer++
		nonZeroPointer++
	}
}