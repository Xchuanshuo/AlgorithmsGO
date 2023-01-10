package test778

var dirs = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func swimInWater(grid [][]int) int {
	var n = len(grid)
	var idxs = make([]int, n*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			idxs[grid[i][j]] = i*n + j
		}
	}
	var uf = NewDefUF()
	for i := 0; i < n*n; i++ {
		x, y := idxs[i]/n, idxs[i]%n
		for j := 0; j < len(dirs); j++ {
			nx, ny := x+dirs[j][0], y+dirs[j][1]
			if nx < 0 || nx >= n || ny < 0 || ny >= n || grid[nx][ny] > i {
				continue
			}
			uf.union(x*n+y, nx*n+ny)
		}
		if uf.isConnected(0, n*n-1) {
			return i
		}
	}
	return n * n
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
