package lcw386

/**
 * @Description https://leetcode.cn/contest/weekly-contest-386/problems/earliest-second-to-mark-indices-i/
 * idea: 二分 + 逆序贪心，栈顶的优先考虑
 **/

func earliestSecondToMarkIndices(nums []int, changeIndices []int) int {
	var n = len(nums)
	var m = len(changeIndices)
	var cal = func(i int) bool {
		var t = make([]int, n)
		var in = make(map[int]bool)
		copy(t, nums)
		var s = make([]int, 0)
		for j := i; j >= 0; j-- {
			var numP = changeIndices[j]
			// 最后一次出现的位置，且当前为0
			if !in[numP] {
				if t[numP-1] > 0 {
					s = append(s, numP)
				}
				in[numP] = true
			} else {
				if len(s) > 0 {
					t[s[len(s)-1]-1]--
					if t[s[len(s)-1]-1] <= 0 {
						s = s[0 : len(s)-1]
					}
				}
			}
		}
		for _, v := range t {
			if v > 0 {
				return false
			}
		}
		return len(in) == n
	}
	l, r := 0, m-1
	for l <= r {
		var mid = l + (r-l)/2
		if !cal(mid) {
			l = mid + 1
		} else {
			if mid == 0 || !cal(mid-1) {
				return mid + 1
			}
			r = mid - 1
		}
	}
	return -1
}
