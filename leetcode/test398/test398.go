package test398

import "math/rand"

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums: nums}
}

func (this *Solution) Pick(target int) int {
	j, value := 0, 0
	var nums = this.nums
	for i := 0; i < len(nums); i++ {
		if nums[i] == target && rand.Intn(j+1) == 0 {
			value = i
			j += 1
		}
	}
	return value
}
