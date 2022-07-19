package test1609


func isEvenOddTree(root *TreeNode) bool {
	q := make([]*TreeNode, 0)
	q = append(q, root)
	level, INF := 0, int(1e6+1)
	for len(q) > 0 {
		cnt := len(q)
		last := 0
		if level%2 == 1 { last = INF }
		for i := 0;i < cnt;i++ {
			cur := q[0]
			if cur.Left != nil { q = append(q, cur.Left) }
			if cur.Right != nil { q = append(q, cur.Right) }
			if level%2 == 0 {
				if cur.Val%2 == 0 || cur.Val <= last { return false }
			} else {
				if cur.Val%2 != 0 || cur.Val >= last { return false }
			}
			last = cur.Val
			q = q[1:]
		}
		level++
	}
	return true
}

type TreeNode struct {
	Val int
	Left, Right *TreeNode
}