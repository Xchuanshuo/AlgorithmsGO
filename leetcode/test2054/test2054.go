package test2054

import (
	"github.com/emirpasic/gods/trees/redblacktree"
)

import "sort"

func maxTwoEvents(events [][]int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][1] < events[j][1]
	})
	var dp = redblacktree.NewWithIntComparator()
	var res = 0
	for _, e := range events {
		node, found := dp.Floor(e[0] - 1)
		var lv = 0
		if found {
			lv = node.Value.(int)
		}
		res = max(res, lv+e[2])
		var cur = e[2]
		node1, found := dp.Floor(e[1])
		if found {
			cur = max(cur, node1.Value.(int))
		}
		dp.Put(e[1], cur)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
