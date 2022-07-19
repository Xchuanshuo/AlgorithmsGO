package test506

import (
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	sMap := make(map[int]int)
	for idx, v := range score {
		sMap[v] = idx
	}
	sort.Ints(score)
	result := make([]string, len(score))
	for i, _ := range score {
		idx := len(score) - i - 1
		cur := score[idx]
		if i == 0 {
			result[sMap[cur]] = "Gold Medal"
		} else if i == 1 {
			result[sMap[cur]] = "Silver Medal"
		} else if i == 2 {
			result[sMap[cur]] = "Bronze Medal"
		} else {
			result[sMap[cur]] = strconv.Itoa(i + 1)
		}
	}
	return result
}