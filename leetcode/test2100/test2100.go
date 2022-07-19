package test2100

func goodDaysToRobBank(security []int, time int) []int {
	var n = len(security)
	var pre = make([]int, n)
	var after = make([]int, n)
	for i := 0; i < len(security); i++ {
		if i > 0 && security[i] <= security[i-1] {
			pre[i] = pre[i-1] + 1
		}
	}
	for i := len(security) - 1; i >= 0; i-- {
		if i+1 < len(security) && security[i] <= security[i+1] {
			after[i] = after[i+1] + 1
		}
	}
	var res []int
	for i := 0; i < len(security); i++ {
		if pre[i] >= time && after[i] >= time {
			res = append(res, i)
		}
	}
	return res
}
