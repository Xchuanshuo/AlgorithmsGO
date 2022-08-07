package test449

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	str string
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var res = serial(root)
	this.str = res
	return res
}

func serial(root *TreeNode) string {
	if root == nil {
		return "#,"
	}
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%s,", strconv.Itoa(root.Val)))
	sb.WriteString(fmt.Sprintf("%s", serial(root.Left)))
	sb.WriteString(fmt.Sprintf("%s", serial(root.Right)))
	return sb.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	q = strings.Split(data, ",")
	return recov()
}

var q []string

func recov() *TreeNode {
	if len(q) == 0 {
		return nil
	}
	var cur = q[0]
	q = q[1:]
	if cur == "#" {
		return nil
	}
	v, _ := strconv.Atoi(cur)
	var root = &TreeNode{Val: v}
	root.Left = recov()
	root.Right = recov()
	return root
}
