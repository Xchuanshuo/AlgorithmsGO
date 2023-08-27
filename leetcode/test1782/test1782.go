package test1782

/**
 * @Description https://leetcode.cn/problems/count-pairs-of-nodes/
 * idea: 哈希表 + 二分查找
		1.统计重复边，建图(不包含重复边)，统计节点的度
		2.对节点的度进行排序，并记录排序后节点的最终位置
		3.对于每个询问，从排序后的每个节点a，二分查找后续的节点b使得 degree[a]+degree[b] > q,
		  即查找左边界l，对于n-l范围内的节点都可以与节点a进行配对，满足数量>q，由于相邻的节点存在【重复边】，所以
		  需要对a的每个在范围[l,n)内的相邻节点b减去重复边后进行二次判定。由于边的数量为1e5，每个边最多被访问两次，
		  时间复杂度满足要求，整体时间为 O(len(e)+len(q)*n*log(n))
 **/

import "sort"

func countPairs(n int, edges [][]int, queries []int) []int {
	var g = make(map[int][]int)
	var mp = make(map[t]int)
	var degree = make([]int, n+1)
	for _, e := range edges {
		v, w := e[0], e[1]
		if v > w {
			v, w = w, v
		}
		var p = t{v, w}
		if mp[p] == 0 {
			g[e[0]] = append(g[e[0]], e[1])
			g[e[1]] = append(g[e[1]], e[0])
		}
		mp[p]++
		degree[e[0]]++
		degree[e[1]]++
	}
	var dv = make([][]int, 0)
	for i, v := range degree {
		dv = append(dv, []int{i, v})
	}
	dv = dv[1:]
	sort.Slice(dv, func(i, j int) bool {
		return dv[i][1] < dv[j][1]
	})
	var pos = make(map[int]int) // 记录排序后每个节点的位置
	for i, v := range dv {
		pos[v[0]] = i
	}
	var res = make([]int, len(queries))
	for i, q := range queries {
		var cur = 0
		for j, d := range dv {
			cntA, a := d[1], d[0]
			var l = sort.Search(len(dv), func(i int) bool {
				return cntA+dv[i][1] > q
			})
			if l <= j {
				l = j + 1
			}
			cur += n - l
			for _, adj := range g[a] {
				if pos[adj] < l {
					continue
				}
				var cntB = degree[adj]
				v, w := a, adj
				if v > w {
					v, w = w, v
				}
				var cnt = cntA + cntB - mp[t{v, w}]
				if cnt <= q { // 二次判定
					cur--
				}
			}
		}
		res[i] = cur
	}
	return res
}

type t struct {
	v, w int
}
