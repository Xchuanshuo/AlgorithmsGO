package test803

/**
 * @Description https://leetcode.cn/problems/bricks-falling-when-hit/
 * idea: 并查集, 计算因为消除操作掉落的砖块 => 添加该砖块后联通分量的大小变化了多少
	   具体步骤: 1.预处理 把原先消除砖块的地方先进行去除，分别合并各个联通分量
			   2.逆序添加消除的砖块，进行新的节点合并，过程中计算结果
			   3.对于稳定的砖块，可以统一合并到唯一根节点-1
 **/

var dirs = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func hitBricks(grid [][]int, hits [][]int) []int {
	m, n := len(grid), len(grid[0])
	var hitMap = make(map[int][]int)
	for idx, h := range hits {
		hitMap[h[0]*n+h[1]] = []int{idx, 1}
	}
	var res = make([]int, len(hits))
	// 1.预处理
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			var pos = i*n + j
			if _, exist := hitMap[pos]; !exist {
				continue
			}
			if hitMap[pos][1] == 1 && grid[i][j] == 1 {
				res[hitMap[pos][0]] = -1
				grid[i][j] = 0
			}
		}
	}
	var root = -1
	var uf = NewDefUF()
	// 2.消除后 节点合并
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			if i == 0 {
				uf.union(root, i*n+j)
			}
			for k := 0; k < len(dirs); k++ {
				nx, ny := i+dirs[k][0], j+dirs[k][1]
				if nx < 0 || nx >= m || ny < 0 || ny >= n || grid[nx][ny] == 0 {
					continue
				}
				uf.union(i*n+j, nx*n+ny)
			}
		}
	}
	for i := len(hits) - 1; i >= 0; i-- {
		x, y := hits[i][0], hits[i][1]
		var pos = x*n + y
		var cur = res[i]
		if cur == 0 {
			continue
		}
		var preSize = uf.sz[uf.find(root)]
		_ = uf.find(pos)
		if x == 0 {
			uf.union(pos, root)
		}
		grid[x][y] = 1
		for k := 0; k < len(dirs); k++ {
			nx, ny := x+dirs[k][0], y+dirs[k][1]
			if nx < 0 || nx >= m || ny < 0 || ny >= n || grid[nx][ny] == 0 {
				continue
			}
			uf.union(pos, nx*n+ny)
		}
		var curSize = uf.sz[uf.find(root)]
		res[i] = max(0, curSize-preSize-1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type UF struct {
	parent map[int]int
	sz     map[int]int
}

func NewUF(arr []int) *UF {
	var parent = make(map[int]int)
	var sz = make(map[int]int)
	for i := 0; i < len(arr); i++ {
		parent[i] = i
		sz[i] = 1
	}
	var uf = &UF{parent: parent, sz: sz}
	return uf
}

func NewDefUF() *UF {
	return NewUF([]int{})
}

func (uf *UF) union(p, q int) {
	var pRoot = uf.find(p)
	var qRoot = uf.find(q)
	var sz = uf.sz
	if pRoot == qRoot {
		return
	}
	if sz[pRoot] > sz[qRoot] {
		sz[pRoot] += sz[qRoot]
		uf.parent[qRoot] = pRoot
	} else {
		sz[qRoot] += sz[pRoot]
		uf.parent[pRoot] = qRoot
	}
}

func (uf *UF) isConnected(p, q int) bool {
	return uf.find(p) == uf.find(q)
}

func (uf *UF) find(p int) int {
	if _, exist := uf.parent[p]; !exist {
		uf.parent[p] = p
		uf.sz[p] = 1
	}
	var parent = uf.parent
	for parent[p] != p {
		parent[p] = parent[parent[p]]
		p = parent[p]
	}
	return p
}
