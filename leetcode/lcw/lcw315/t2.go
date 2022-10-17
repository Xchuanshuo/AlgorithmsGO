package lcw315

import (
	"strconv"
)

func countDistinctIntegers(nums []int) int {
	var mp = make(map[int]bool)
	for _, num := range nums {
		mp[num] = true
		var s = []byte(strconv.Itoa(num))
		for j := 0; j < len(s)/2; j++ {
			s[j], s[len(s)-j-1] = s[len(s)-j-1], s[j]
		}
		v, _ := strconv.Atoi(string(s))
		mp[v] = true
	}
	return len(mp)
}
