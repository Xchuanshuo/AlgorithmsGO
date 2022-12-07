package lcw319

import (
	"sort"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minimumOperations(root *TreeNode) int {
	var q = make([]*TreeNode, 0)
	q = append(q, root)
	var level = 0
	var levelOrders = make(map[int]map[int]int)
	for len(q) > 0 {
		var size = len(q)
		var arr = make([]int, 0)
		for i := 0; i < size; i++ {
			var cur = q[0]
			q = q[1:]
			arr = append(arr, cur.Val)
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		var orders = make(map[int]int)
		sort.Ints(arr)
		for i, v := range arr {
			orders[i] = v
		}
		levelOrders[level] = orders
		level++
	}
	level, res := 0, 0
	q = append(q, root)
	for len(q) > 0 {
		var size = len(q)
		var arr = make([]int, 0)
		var curOrders = make(map[int]int)
		for i := 0; i < size; i++ {
			var cur = q[0]
			q = q[1:]
			arr = append(arr, cur.Val)
			curOrders[cur.Val] = i
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		var orders = levelOrders[level]
		for i, v := range arr {
			if v == orders[i] {
				continue
			}
			var idx = curOrders[orders[i]]
			curOrders[arr[i]], curOrders[arr[idx]] = curOrders[arr[idx]], curOrders[arr[i]]
			arr[i], arr[idx] = arr[idx], arr[i]
			res++
		}
		level++
	}
	return res
}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}
