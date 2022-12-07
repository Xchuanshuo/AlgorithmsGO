package main

import (
	"sort"
	"strings"
)

// 1.找总量最大 2.最大中播放量最多
func mostPopularCreator(creators []string, ids []string, views []int) [][]string {
	var mp = make(map[string][]Pair)
	var sum = make(map[string]int)
	var mx = 0
	for i, c := range creators {
		mp[c] = append(mp[c], Pair{ids[i], views[i]})
		sum[c] += views[i]
		if sum[c] > mx {
			mx = sum[c]
		}
	}
	var res = make([][]string, 0)
	for k, pairs := range mp {
		if sum[k] == mx {
			sort.Slice(pairs, func(i, j int) bool {
				if pairs[i].view != pairs[j].view {
					return strings.Compare(pairs[i].id, pairs[j].id) <= 0
				}
				return pairs[i].view > pairs[j].view
			})
			res = append(res, []string{k, pairs[0].id})
		}
	}
	return res
}

type Pair struct {
	id   string
	view int
}
