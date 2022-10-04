package main

import "sort"

/**
 * @Description https://leetcode.cn/problems/number-of-good-paths/
 * idea: 关键条件【图中所有节点联通，树形结构 两点之间路径唯一】 使用并查集
 *       1.节点值从小到大排序 按此顺序进行合并
 *		 2.合并时将值较小的节点合并到较大的节点
 * 		 3.计算好路径需要看两个联通分量【最大值节点】的个数 使用map存储
 *		 4.好路径数量为两个联通分量最大值个数相乘
 * 参考: https://www.bilibili.com/video/BV1ve411K7P5/?vd_source=4ace28b4fa14f8b7a8597a27330243a1
 **/

func numberOfGoodPaths(vals []int, edges [][]int) int {
	var n = len(vals)
	var graph = make([][]int, n)
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}
	var parent = make([]int, n)
	var find func(p int) int
	find = func(id int) int {
		for id != parent[id] {
			id = parent[id]
		}
		return parent[id]
	}
	var mp = make(map[int]int)
	var ids = make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		mp[i] = 1
		ids[i] = i
	}
	sort.Slice(ids, func(i, j int) bool {
		return vals[ids[i]] < vals[ids[j]]
	})
	var res = n
	for i := 0; i < n; i++ {
		pv, pRoot := vals[ids[i]], find(ids[i])
		for _, nxt := range graph[ids[i]] {
			var qRoot = find(nxt)
			var qv = vals[qRoot]
			// 只合并比当前节点小的节点
			if qv > pv || pRoot == qRoot {
				continue
			}
			if qv == pv { // 满足最好路径
				res += mp[pRoot] * mp[qRoot]
				mp[pRoot] += mp[qRoot]
			}
			parent[qRoot] = pRoot
		}
	}
	return res
}
