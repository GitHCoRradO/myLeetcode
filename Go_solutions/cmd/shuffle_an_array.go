package main

import (
	"math/rand"
)
type solution struct {
	nums []int
	original []int
}


func Constructor(nums []int) solution {
	return solution{nums, append([]int(nil), nums...)}
}


func (this *solution) Reset() []int {
	copy(this.nums, this.original)
	return this.nums
}


func (this *solution) Shuffle() []int {
	//rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(this.nums), func(i, j int) {
		this.nums[i], this.nums[j] = this.nums[j], this.nums[i]
	})
	return this.nums
}


/**
 * Your solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
