package test1728

import "sync"

const K = 1000

var dirs = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

const State = 4096

var dp = make([][]int, State)
var cj, mj int
var fx, fy int
var m, n int
var g []string

func canMouseWin(grid []string, catJump int, mouseJump int) bool {
	cj, mj = catJump, mouseJump
	g = grid
	var wg = sync.WaitGroup{}
	var cnt = 8
	wg.Add(cnt)
	var f = func(start, end int) {
		for i := start; i < end; i++ {
			if dp[i] == nil {
				dp[i] = make([]int, K)
			}
			for j := 0; j < K; j++ {
				dp[i][j] = -1
			}
		}
		wg.Done()
	}
	var delta = State / cnt
	for i := 0; i < cnt; i++ {
		go f(i*delta, (i+1)*delta)
	}
	wg.Wait()
	cx, cy, mx, my := 0, 0, 0, 0
	m, n = len(g), len(g[0])
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 'C' {
				cx, cy = i, j
			} else if grid[i][j] == 'M' {
				mx, my = i, j
			} else if grid[i][j] == 'F' {
				fx, fy = i, j
			}
		}
	}
	return dfs(cx, cy, mx, my, 0) == 0
}

func dfs(cx, cy, mx, my, k int) int {
	if k+1 == K {
		return 1
	}
	if cx == mx && cy == my {
		return 1
	}
	if cx == fx && cy == fy {
		return 1
	}
	if mx == fx && my == fy {
		return 0
	}
	var state = cx<<9 | cy<<6 | mx<<3 | my
	if dp[state][k] != -1 {
		return dp[state][k]
	}
	if k%2 == 0 { // mouse
		for j := 0; j < len(dirs); j++ {
			for i := 0; i <= mj; i++ {
				nx, ny := dirs[j][0]*i+mx, dirs[j][1]*i+my
				if nx < 0 || nx >= m || ny < 0 || ny >= n || g[nx][ny] == '#' {
					break
				}
				var res = dfs(cx, cy, nx, ny, k+1)
				if res == 0 {
					dp[state][k] = 0
					return 0
				}
			}
		}
		dp[state][k] = 1
	} else { // cat
		for j := 0; j < len(dirs); j++ {
			for i := 0; i <= cj; i++ {
				nx, ny := dirs[j][0]*i+cx, dirs[j][1]*i+cy
				if nx < 0 || nx >= m || ny < 0 || ny >= n || g[nx][ny] == '#' {
					break
				}
				var res = dfs(nx, ny, mx, my, k+1)
				if res == 1 {
					dp[state][k] = 1
					return 1
				}
			}
		}
		dp[state][k] = 0
	}
	return dp[state][k]
}
