package lcw385

import "strconv"

func longestCommonPrefix(arr1 []int, arr2 []int) int {
	for l := 9; l >= 1; l-- {
		var set = make(map[string]bool)
		for _, v := range arr1 {
			var s = strconv.Itoa(v)
			if len(s) < l {
				continue
			}
			set[s[0:l]] = true
		}
		for _, v := range arr2 {
			var s = strconv.Itoa(v)
			if len(s) < l {
				continue
			}
			if set[s[0:l]] {
				return l
			}
		}
	}
	return 0
}
