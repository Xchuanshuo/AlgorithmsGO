package test652

import (
	"fmt"
	"strconv"
)

var mp map[string]*TreeNode

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	mp = make(map[string]*TreeNode)
	dfs(root)
	var res = make([]*TreeNode, 0)
	for _, v := range mp {
		if v == nil {
			continue
		}
		res = append(res, v)
	}
	return res
}

func dfs(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	var left = fmt.Sprintf("l%s", dfs(root.Left))
	var right = fmt.Sprintf("r%s", dfs(root.Right))
	var cur = fmt.Sprintf("%s%s%s", strconv.Itoa(root.Val), left, right)
	if _, exist := mp[cur]; exist {
		mp[cur] = root
	} else {
		mp[cur] = nil
	}
	return cur
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
