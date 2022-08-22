package test655

import (
	"math"
	"strconv"
)

var grid [][]string
var h, n int

func printTree(root *TreeNode) [][]string {
	if root == nil {
		return [][]string{}
	}
	h = getH(root) - 1
	n = pow(2, h+1) - 1
	grid = make([][]string, h+1)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]string, n)
	}
	var mid = (n - 1) / 2
	grid[0][mid] = strconv.Itoa(root.Val)
	dfs(root.Left, 1, mid, -pow(2, h-1))
	dfs(root.Right, 1, mid, pow(2, h-1))
	return grid
}

func dfs(root *TreeNode, r, c int, pos int) {
	if root == nil {
		return
	}
	var nc = c + pos
	grid[r][nc] = strconv.Itoa(root.Val)
	dfs(root.Left, r+1, nc, -pow(2, h-r-1))
	dfs(root.Right, r+1, nc, pow(2, h-r-1))
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func getH(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(getH(root.Left), getH(root.Right))
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
