package test_off2_29

func insert(aNode *Node, x int) *Node {
	if aNode == nil {
		var newNode = &Node{Val: x}
		newNode.Next = newNode
		return newNode
	}
	var h = aNode
	var max *Node = nil
	for h != nil {
		if max == nil || max.Val <= h.Val {
			max = h
		}
		if x >= h.Val && x <= h.Next.Val {
			var newNode = &Node{Val: x}
			newNode.Next = h.Next
			h.Next = newNode
			break
		}
		h = h.Next
		if h == aNode {
			var newNode = &Node{Val: x}
			newNode.Next = max.Next
			max.Next = newNode
			break
		}
	}
	return aNode
}

type Node struct {
	Val  int
	Next *Node
}
