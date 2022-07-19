package test515

func largestValues(root *TreeNode) []int {
	var res = make([]int, 0)
	if root == nil {
		return res
	}
	var q = make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) > 0 {
		var max *TreeNode = nil
		var size = len(q)
		for i := 0; i < size; i++ {
			var cur = q[i]
			if max == nil || cur.Val > max.Val {
				max = cur
			}
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		q = q[size:]
		if max != nil {
			res = append(res, max.Val)
		}
	}
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
