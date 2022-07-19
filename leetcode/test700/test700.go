package test700

import "LeetCodeGO"

func searchBST(root *AlgorithmsGO.TreeNode, val int) *AlgorithmsGO.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > val {
		return searchBST(root.Left, val)
	} else if root.Val < val {
		return searchBST(root.Right, val)
	} else {
		return root
	}
}
