package test590

func postorder1(root *Node) []int {
	type tNode struct {
		node *Node
		cnt  int
	}
	var s = make([]*tNode, 0)
	var res = make([]int, 0)
	if root == nil { return res }
	s = append(s, &tNode{node: root, cnt: 0})
	for len(s) > 0 {
		cur := s[len(s)-1]
		s = s[0 : len(s)-1]
		node, cnt := cur.node, cur.cnt
		if cnt == len(node.Children) {
			res = append(res, node.Val)
		}
		if cnt < len(node.Children) {
			s = append(s, &tNode{node: node, cnt: cnt + 1})
			s = append(s, &tNode{node: node.Children[cnt], cnt: 0})
		}
	}
	return res
}