package lcw311

//[2,3,5,7,11,13,17,19,23,29,31,37,41,43,47]
//输出：
//[2,5,3,7,11,13,17,37,41,43,47,19,23,29,31]
//预期：
//[2,5,3,7,11,13,17,47,43,41,37,31,29,23,19]

func reverseOddLevels(root *TreeNode) *TreeNode {
	var q = make([]*TreeNode, 0)
	q = append(q, root)
	var level = 0
	for len(q) > 0 {
		var sz = len(q)
		var tmp = make([]*TreeNode, 0)
		for i := 0; i < sz; i++ {
			var cur = q[0]
			q = q[1:]
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
			tmp = append(tmp, cur)
		}
		if level%2 == 1 {
			for i := 0; i < sz/2; i++ {
				tmp[i].Val, tmp[sz-i-1].Val = tmp[sz-i-1].Val, tmp[i].Val
			}
		}
		level++
	}
	return root
}

func dfs(root *TreeNode, level int) {
	if root == nil {
		return
	}
	dfs(root.Left, level+1)
	dfs(root.Right, level+1)
	if level%2 == 1 {
		if root.Left == nil && root.Right == nil {
			return
		}
		root.Left.Val, root.Right.Val = root.Right.Val, root.Left.Val
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
