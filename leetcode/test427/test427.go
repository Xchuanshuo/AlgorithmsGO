package test427

func construct(grid [][]int) *Node {
	var n = len(grid)
	var sums = make([][]int, n+1)
	for i := 0; i < len(sums); i++ {
		sums[i] = make([]int, n+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sums[i+1][j+1] = sums[i][j+1] + sums[i+1][j] - sums[i][j] + grid[i][j]
		}
	}
	var dfs func(x1, y1, x2, y2 int) *Node
	dfs = func(x1, y1, x2, y2 int) *Node {
		var val = sums[x2][y2] - sums[x2][y1] - sums[x1][y2] + sums[x1][y1]
		if val == 0 {
			return NewNode(false, true, nil, nil, nil, nil)
		}
		if val == (x2-x1)*(y2-y1) {
			return NewNode(true, true, nil, nil, nil, nil)
		}
		xm, ym := (x2+x1)/2, (y2+y1)/2
		return NewNode(true, false,
			dfs(x1, y1, xm, ym),
			dfs(x1, ym, xm, y2),
			dfs(xm, y1, x2, ym),
			dfs(xm, ym, x2, y2))
	}
	return dfs(0, 0, n, n)
}

func NewNode(val, isLeaf bool, tLeft, tRight, bLeft, bRight *Node) *Node {
	return &Node{Val: val, IsLeaf: isLeaf, TopLeft: tLeft, TopRight: tRight,
		BottomLeft: bLeft, BottomRight: bRight}
}

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}
