package test590

var res []int

func postorder(root *Node) []int {
	res = make([]int, 0)
	solve(root)
	return res
}

func solve(root *Node) {
	if root == nil {
		return
	}
	for _, c := range root.Children {
		solve(c)
	}
	res = append(res, root.Val)
}

type Node struct {
	Val      int
	Children []*Node
}
