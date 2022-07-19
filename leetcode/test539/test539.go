package test539

import "sort"

func findMinDifference(times []string) int {
	sort.Strings(times)
	min, pre := 1440, getMin(times[0])
	for i := 1;i < len(times);i++ {
		cur := getMin(times[i])
		dif := cur - pre
		if dif < min { min = dif }
		pre = cur
	}
	v := getMin(times[0]) + 1440 - pre
	if v < min { min = v }
	return min
}

func getMin(s string) int {
	return int((s[0]-'0')*10+ s[1]-'0')*60 + int((s[3]-'0')*10 + s[4]-'0')
}