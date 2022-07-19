package test563

import "LeetCodeGO"

var res int
func findTilt(root *AlgorithmsGO.TreeNode) int {
	res = 0
	dfs(root)
	return res
}

func dfs(root *AlgorithmsGO.TreeNode) int {
	if root == nil { return 0}
	left := dfs(root.Left)
	right := dfs(root.Right)
	res += abs(left, right)
	return root.Val + left + right
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
