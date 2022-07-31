package test1161

import "math"

func maxLevelSum(root *TreeNode) int {
	var q = make([]*TreeNode, 0)
	q = append(q, root)
	sum, res := math.MinInt32, 0
	var layer = 0
	for len(q) > 0 {
		layer++
		var size = len(q)
		var cur = 0
		for i := 0; i < size; i++ {
			var c = q[0]
			q = q[1:]
			cur += c.Val
			if c.Left != nil {
				q = append(q, c.Left)
			}
			if c.Right != nil {
				q = append(q, c.Right)
			}
		}
		if cur > sum {
			sum = cur
			res = layer
		}
	}
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
