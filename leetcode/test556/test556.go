package test556

import (
	"math"
	"strconv"
)

func nextGreaterElement(n int) int {
	var str = strconv.Itoa(n)
	var nums = []byte(str)
	n = len(nums)
	var i = n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i == -1 {
		return -1
	}
	var pos = n - 1
	for nums[pos] <= nums[i] {
		pos--
	}
	nums[i], nums[pos] = nums[pos], nums[i]
	i += 1
	n -= 1
	for i <= n {
		nums[i], nums[n] = nums[n], nums[i]
		i += 1
		n -= 1
	}
	val, _ := strconv.ParseInt(string(nums), 10, 64)
	if val > math.MaxInt32 {
		return -1
	}
	return int(val)
}
