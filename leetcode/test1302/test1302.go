package test1302

var maxLevel, sum int

func deepestLeavesSum(root *TreeNode) int {
	maxLevel, sum = -1, 0
	dfs(root, 0)
	return sum
}

func dfs(root *TreeNode, level int) {
	if root == nil {
		return
	}
	if level > maxLevel {
		maxLevel = level
		sum = root.Val
	} else if level == maxLevel {
		sum += root.Val
	}
	dfs(root.Left, level+1)
	dfs(root.Right, level+1)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
