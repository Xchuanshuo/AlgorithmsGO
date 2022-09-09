package test998

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val > root.Val {
		var newRoot = &TreeNode{Val: val}
		newRoot.Left = root
		return newRoot
	}
	root.Right = insertIntoMaxTree(root.Right, val)
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
