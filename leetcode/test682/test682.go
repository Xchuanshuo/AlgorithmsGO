package test682

import "strconv"

func calPoints(ops []string) int {
	var arr = make([]int, len(ops))
	var idx = -1
	for _, v := range ops {
		if v == "C" {
			arr[idx] = 0
			idx--
		} else if v == "D" {
			var pre = arr[idx]
			idx++
			arr[idx] = 2 * pre
		} else if v == "+" {
			idx++
			arr[idx] = arr[idx-1] + arr[idx-2]
		} else {
			idx++
			val, _ := strconv.Atoi(v)
			arr[idx] = val
		}
	}
	var res = 0
	for i := 0; i < len(arr); i++ {
		res += arr[i]
	}
	return res
}
