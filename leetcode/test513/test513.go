package test513

func findBottomLeftValue(root *TreeNode) int {
	var q = make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) > 0 {
		root = q[0]
		q = q[1:]
		if root.Right != nil {
			q = append(q, root.Right)
		}
		if root.Left != nil {
			q = append(q, root.Left)
		}
	}
	return root.Val
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
