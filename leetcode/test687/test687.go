package test687

var res = 0

func longestUnivaluePath(root *TreeNode) int {
	res = 0
	dfs(root)
	return res
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var rtnL, rtnR int
	var lv = dfs(root.Left)
	var rv = dfs(root.Right)
	if root.Left != nil && root.Val == root.Left.Val {
		rtnL = lv
	}
	if root.Right != nil && root.Val == root.Right.Val {
		rtnR = rv
	}
	res = max(res, 1+rtnL+rtnR)
	return 1 + max(rtnL, rtnR)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
