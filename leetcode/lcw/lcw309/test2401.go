package lcw309

/**
 * @Description https://leetcode.cn/contest/weekly-contest-309/problems/longest-nice-subarray/
 * idea: 滑动窗口 范围内的二进制位为1只能出现一次
 **/

func longestNiceSubarray(nums []int) int {
	var n = len(nums)
	l, r := 0, 0
	var wind = make([]int, 32)
	var cal = func(a int, isInc bool) bool {
		var isMove = false
		for i := 0; i < 32; i++ {
			if (1<<i)&a > 0 {
				if isInc {
					wind[i]++
				} else {
					wind[i]--
				}
			}
			if wind[i] > 1 {
				isMove = true
			}
		}
		return isMove
	}
	var res = 1
	for r < n {
		var cur = nums[r]
		isMove := cal(cur, true)
		for isMove && l <= r {
			cur = nums[l]
			isMove = cal(cur, false)
			l++
		}
		res = max(res, r-l+1)
		r++
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
