package test606

import (
	"strconv"
	"strings"
)

var sb strings.Builder

func tree2str(root *TreeNode) string {
	sb.Reset()
	preOrder(root)
	return sb.String()
}

func preOrder(root *TreeNode) {
	if root == nil {
		return
	}
	sb.WriteString(strconv.Itoa(root.Val))
	if root.Left != nil {
		sb.WriteString("(")
		preOrder(root.Left)
		sb.WriteString(")")
	}
	if root.Right != nil {
		if root.Left == nil {
			sb.WriteString("()")
		}
		sb.WriteString("(")
		preOrder(root.Right)
		sb.WriteString(")")
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
