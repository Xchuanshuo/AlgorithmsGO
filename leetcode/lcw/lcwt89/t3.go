package main

import "fmt"

func minimizeArrayValue(nums []int) int {
	n, sum := len(nums), 0
	var check = func(t int) bool {
		var cnt = 0
		for i := n - 1; i >= 0; i-- {
			var cur = nums[i]
			cur += cnt
			if cur > t {
				cnt = cur - t
			} else {
				cnt = 0
			}
		}
		return cnt == 0
	}
	for _, num := range nums {
		sum += num
	}
	l, r := 0, sum
	for l <= r {
		mid := l + (r-l)/2
		if !check(mid) {
			l = mid + 1
		} else {
			if mid == 0 || !check(mid-1) {
				return mid
			}
			r = mid - 1
		}
	}
	return 0
}

func main() {
	var arr = []int{3, 7, 1, 6}
	res := minimizeArrayValue(arr)
	fmt.Println(res)
}
