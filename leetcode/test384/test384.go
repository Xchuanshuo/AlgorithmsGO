package test384

import "math/rand"

type Solution struct {
	nums []int
}


func Constructor(nums []int) Solution {
	return Solution{
		nums: nums,
	}
}


func (this *Solution) Reset() []int {
	return this.nums
}


func (this *Solution) Shuffle() []int {
	n := len(this.nums)
	var clone = make([]int, n)
	copy(clone, this.nums)
	for i := 0;i < n;i++ {
		r := i + rand.Intn(n - i)
		clone[i], clone[r] = clone[r], clone[i]
	}
	return clone
}