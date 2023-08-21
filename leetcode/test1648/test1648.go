package test1648

/**
 * @Description https://leetcode.cn/problems/sell-diminishing-valued-colored-balls/
 * idea: 枚举最多剩remain个球时总共能取的球数，剩余越多，价值越大，呈单调性，所以可以进行二分
	   注意点: remain是下界，有可能刚好取超过remain的球才达到orders个，也可能还需要从剩余的remain个里取，
			  直接按remain取可能导致结果出错
			  所以计算时可以分两步，1.先取超过remain的部分 2.有剩余才从多组remain里面取1个
 **/

var M = int64(1e9) + 7

func maxProfit(invs []int, orders int) int {
	var calc = func(l, r int64) int64 {
		return (l + r) * (r - l + 1) / 2
	}
	var check = func(remain int) bool {
		var res = 0
		for _, v := range invs {
			if v <= remain {
				continue
			}
			res += v - remain
			if res >= orders {
				return true
			}
		}
		return false
	}
	l, r := 0, int(1e9)
	var mid = 0
	for l <= r {
		mid = l + (r-l)/2
		if !check(mid) {
			r = mid - 1
		} else {
			if mid == int(1e9) || !check(mid+1) {
				break
			}
			l = mid + 1
		}
	}
	mid += 1
	var val int64 = 0
	for _, v := range invs {
		if v <= mid {
			continue
		}
		var cnt = min(v-mid, orders)
		val = (val + calc(int64(v-cnt+1), int64(v))) % M
		orders -= cnt
	}
	val = (val + int64(mid*orders)) % M
	return int(val)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
