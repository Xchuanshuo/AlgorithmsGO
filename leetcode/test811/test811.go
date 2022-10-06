package test811

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	var cntMap = make(map[string]int)
	for _, cp := range cpdomains {
		var strs1 = strings.Split(cp, " ")
		cnt, _ := strconv.Atoi(strs1[0])
		cntMap[strs1[1]] += cnt
		for i, v := range strs1[1] {
			if v == '.' {
				cntMap[strs1[1][i+1:]] += cnt
			}
		}
	}
	var res = make([]string, 0)
	for k, v := range cntMap {
		res = append(res, fmt.Sprintf("%d %s", v, k))
	}
	return res
}
