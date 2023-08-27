package main

import (
	"fmt"
)

// 位运算 + 记忆化搜索

var dirs = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func minFlips(g [][]int) int {
	m, n := len(g), len(g[0])
	var src = 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			src |= g[i][j] << (i*n + j)
		}
	}
	var INF = int(1e9)
	var mem = make(map[int]int)
	var dfs func(state int, vis int) int
	dfs = func(state int, vis int) int {
		if state == 0 {
			return 0
		}
		if v, exist := mem[state]; exist {
			return v
		}
		var res = INF
		for i := 0; i < m*n; i++ {
			if vis&(1<<i) != 0 {
				continue
			}
			var cur = reverseBit(state, i)
			x, y := i/n, i%n
			for _, d := range dirs {
				nx, ny := x+d[0], y+d[1]
				if nx < 0 || nx >= m || ny < 0 || ny >= n {
					continue
				}
				var p = nx*n + ny
				cur = reverseBit(cur, p)
			}
			res = min(res, 1+dfs(cur, vis|1<<i))
		}
		mem[state] = res
		return res
	}
	var res = dfs(src, 0)
	if res >= INF {
		return -1
	}
	return res
}

// 翻转位
func reverseBit(num int, i int) int {
	if num&(1<<i) == 0 {
		num |= 1 << i
	} else {
		num &= ^(1 << i)
	}
	return num
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var mat = [][]int{{0, 0}, {0, 1}}
	var res = minFlips(mat)
	fmt.Println(res)
}
