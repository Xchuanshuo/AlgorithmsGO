package test558

func intersect(quadTree1 *Node, quadTree2 *Node) *Node {
	if quadTree1.IsLeaf {
		if quadTree1.Val {
			return &Node{Val: true, IsLeaf: true}
		}
		return quadTree2
	}
	if quadTree2.IsLeaf {
		return intersect(quadTree2, quadTree1)
	}
	t1 := intersect(quadTree1.TopLeft, quadTree2.TopLeft)
	t2 := intersect(quadTree1.TopRight, quadTree2.TopRight)
	t3 := intersect(quadTree1.BottomLeft, quadTree2.BottomLeft)
	t4 := intersect(quadTree1.BottomRight, quadTree2.BottomRight)
	if t1.IsLeaf && t2.IsLeaf && t3.IsLeaf && t4.IsLeaf &&
		t1.Val == t2.Val && t2.Val == t3.Val && t3.Val == t4.Val {
		return &Node{Val: true, IsLeaf: true}
	}
	return &Node{false, false, t1, t2, t3, t4}
}

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}