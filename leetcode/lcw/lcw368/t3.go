package lcw368

// 时间复杂度计算 外层循环*里层循环 < 原数组长度
func minGroupsForValidAssignment(nums []int) int {
	var mp = make(map[int]int)
	for _, v := range nums {
		mp[v]++
	}
	var g = len(nums)
	for _, v := range mp {
		g = min(g, v)
	}
	var cal = func(g int) (int, bool) {
		if g <= 1 {
			return len(nums), true
		}
		var cnt = 0
		for _, v := range mp {
			if v <= g {
				cnt++
			} else {
				var k = g - 1
				var t = v / k
				if v-t*k > t {
					return len(nums), false
				}
				cnt += (v + g - 1) / g
			}
		}
		return cnt, true
	}
	var res = len(nums)
	for i := g + 1; i >= 1; i-- {
		cnt, valid := cal(i)
		if valid {
			res = min(res, cnt)
		}
	}
	return res
}
