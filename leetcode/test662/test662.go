package test662

func widthOfBinaryTree(root *TreeNode) int {
	var q = make([]*Node, 0)
	var res = 1
	q = append(q, &Node{root, 1})
	for len(q) > 0 {
		size := len(q)
		res = max(res, q[size-1].idx-q[0].idx+1)
		for i := 0; i < size; i++ {
			var p = q[0]
			q = q[1:]
			if p.root.Left != nil {
				q = append(q, &Node{p.root.Left, 2 * p.idx})
			}
			if p.root.Right != nil {
				q = append(q, &Node{p.root.Right, 2*p.idx + 1})
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

type Node struct {
	root *TreeNode
	idx  int
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
