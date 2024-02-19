package lcw371

import (
	"sort"
	"strconv"
)

func findHighAccessEmployees(ats [][]string) []string {
	var cal = func(time string) int {
		v1, _ := strconv.Atoi(time[0:2])
		v2, _ := strconv.Atoi(time[2:4])
		return v1*60 + v2
	}
	var mp = make(map[string][]int)
	for _, at := range ats {
		mp[at[0]] = append(mp[at[0]], cal(at[1]))
	}
	var res = make([]string, 0)
	for k, v := range mp {
		sort.Ints(v)
		var valid = false
		for i := 1; i < len(v)-1; i++ {
			if v[i+1]-v[i-1] < 60 {
				valid = true
				break
			}
		}
		if valid {
			res = append(res, k)
		}
	}
	return res
}
