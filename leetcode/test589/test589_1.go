package test589

func preorder1(root *Node) []int {
	var s = make([]*Node, 0)
	var res = make([]int, 0)
	s = append(s, root)
	for len(s) > 0 {
		cur := s[len(s)-1]
		s = s[0 : len(s)-1]
		if cur == nil {
			continue
		}
		res = append(res, cur.Val)
		if cur.Children == nil {
			continue
		}
		for i := len(cur.Children) - 1; i >= 0; i-- {
			child := cur.Children[i]
			s = append(s, child)
		}
	}
	return res
}
