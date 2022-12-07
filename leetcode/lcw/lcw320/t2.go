package lcw320

import "sort"

func closestNodes(root *TreeNode, queries []int) [][]int {
	var arr = make([]int, 0)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		arr = append(arr, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	var findL func(t int) int
	findL = func(t int) int {
		l, r := 0, len(arr)-1
		for l <= r {
			mid := l + (r-l)/2
			if arr[mid] > t {
				r = mid - 1
			} else {
				if mid == len(arr)-1 || arr[mid+1] > t {
					return mid
				}
				l = mid + 1
			}
		}
		return -1
	}
	var res = make([][]int, 0)
	for _, q := range queries {
		var l = findL(q)
		var r = sort.Search(len(arr), func(i int) bool {
			return arr[i] >= q
		})
		res = append(res, []int{l, r})
	}
	return res
}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}
