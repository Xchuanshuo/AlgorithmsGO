package test589

var res []int

func preorder(root *Node) []int {
	res = make([]int, 0)
	solve(root)
	return res
}

func solve(root *Node) {
	if root == nil {
		return
	}
	res = append(res, root.Val)
	children := root.Children
	if children != nil {
		for i := 0; i < len(children); i++ {
			solve(children[i])
		}
	}
}

type Node struct {
	Val      int
	Children []*Node
}
