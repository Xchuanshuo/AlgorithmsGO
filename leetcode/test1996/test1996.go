package test1996

import "sort"

func numberOfWeakCharacters(properties [][]int) int {
	sort.Slice(properties, func(i, j int) bool {
		if properties[i][0] != properties[j][0] {
			return properties[j][0] < properties[i][0]
		}
		return properties[i][1] < properties[j][1]
	})
	res, maxDef := 0, 0
	for _, p := range properties {
		if p[1] < maxDef {
			res++
		} else {
			maxDef = p[1]
		}
	}
	return res
}