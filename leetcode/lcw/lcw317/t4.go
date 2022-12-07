package main

func treeQueries(root *TreeNode, queries []int) []int {
	var height = make(map[*TreeNode]int)
	var getH func(root *TreeNode) int
	getH = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		var h = 1 + max(getH(root.Left), getH(root.Right))
		height[root] = h
		return h
	}
	var maxMap = make(map[int]int)
	var dfs func(root *TreeNode, depth, maxH int)
	dfs = func(root *TreeNode, depth, maxH int) {
		if root == nil {
			return
		}
		maxMap[root.Val] = maxH
		dfs(root.Left, depth+1, max(maxH, depth+height[root.Right]))
		dfs(root.Right, depth+1, max(maxH, depth+height[root.Left]))
	}
	getH(root)
	dfs(root, 0, 0)
	var res = make([]int, 0)
	for _, q := range queries {
		res = append(res, maxMap[q])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}
