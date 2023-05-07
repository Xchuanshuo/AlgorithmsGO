package lcw340

import (
	"math"
	"sort"
)

// 贪心: 排序后选大小相邻的差值;
// 易错点: 假设针对所有差值选, 选p个.
//		  1.将差值排序, 对于差值相等的数对，无法决策从队列中该淘汰哪一个才不会对候选数对产生影响
//		  2.从小到大处理索引, 要获取p个数对，问题同1
// 正解, 使用二分: 只需要判定所有差值相邻数对，有>=p个即可，无需决策该保留哪些

func minimizeMax(nums []int, p int) int {
	if p == 0 {
		return 0
	}
	sort.Ints(nums)
	var n = len(nums)
	var check = func(t int) bool {
		var cnt = 0
		for i := 1; i < n; {
			if nums[i]-nums[i-1] <= t {
				cnt++
				i += 2
			} else {
				i += 1
			}
			if cnt >= p {
				return true
			}
		}
		return false
	}
	l, r := 0, math.MaxInt32
	for l <= r {
		var mid = l + (r-l)/2
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
