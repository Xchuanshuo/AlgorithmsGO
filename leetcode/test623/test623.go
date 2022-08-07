package test623

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if root == nil {
		return nil
	}
	if depth == 1 {
		return &TreeNode{val, root, nil}
	}
	if depth == 2 {
		root.Left = &TreeNode{val, root.Left, nil}
		root.Right = &TreeNode{val, nil, root.Right}
	} else {
		root.Left = addOneRow(root.Left, val, depth-1)
		root.Right = addOneRow(root.Right, val, depth-1)
	}
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
