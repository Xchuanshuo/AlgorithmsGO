package test1405

import "sort"

func longestDiverseString(a int, b int, c int) string {
	var arr = [][]byte{{'a', byte(a)}, {'b', byte(b)}, {'c', byte(c)}}
	var res = make([]byte, 0)
	for true {
		sort.Slice(arr, func(i, j int) bool {
			return arr[j][1] < arr[i][1]
		})
		var hasChar = false
		for idx, cur := range arr {
			if cur[1] <= 0 {
				break
			}
			l := len(res)
			if l >= 2 && res[l-1] == cur[0] && res[l-2] == cur[0] {
				if idx >= 2 || arr[idx+1][1] <= 0 {
					break
				}
				cur = arr[idx+1]
				cur[1]--
				res = append(res, cur[0])
				hasChar = true
				break
			} else {
				cur[1]--
				res = append(res, cur[0])
				hasChar = true
				break
			}
		}
		if !hasChar {
			return string(res)
		}
	}
	return ""
}
