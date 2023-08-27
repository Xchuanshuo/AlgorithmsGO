package test2096

import "strings"

func getDirections(root *TreeNode, s int, d int) string {
	var path = make([]byte, 0)
	var dfs func(root *TreeNode, t int, to int) bool
	dfs = func(root *TreeNode, t int, to int) bool {
		if t == 1 {
			path = append(path, 'L')
		} else if t == 2 {
			path = append(path, 'R')
		}
		if root == nil {
			return false
		}
		if root.Val == to {
			return true
		}
		var cnt = 0
		if !dfs(root.Left, 1, to) && len(path) > 0 { // 不合法则回溯
			path = path[0 : len(path)-1]
			cnt++
		}
		if !dfs(root.Right, 2, to) && len(path) > 0 { // 不合法则回溯
			path = path[0 : len(path)-1]
			cnt++
		}
		return cnt < 2
	}
	dfs(root, 0, s)
	var path1 = make([]byte, len(path))
	copy(path1, path)

	path = make([]byte, 0)
	dfs(root, 0, d)
	var path2 = make([]byte, len(path))
	copy(path2, path)
	var idx = 0
	for i := 0; i < len(path1) && i < len(path2); i++ {
		if path1[i] != path2[i] {
			break
		}
		idx++
	}
	path1, path2 = path1[idx:], path2[idx:]
	return strings.Repeat("U", len(path1)) + string(path2)
}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}
