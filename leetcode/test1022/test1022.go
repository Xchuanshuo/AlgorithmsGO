package test1022

func sumRootToLeaf(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *TreeNode, val int) int {
	if root == nil {
		return 0
	}
	val = val<<1 | root.Val
	if root.Left == nil && root.Right == nil {
		return val
	}
	return dfs(root.Left, val) + dfs(root.Right, val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
