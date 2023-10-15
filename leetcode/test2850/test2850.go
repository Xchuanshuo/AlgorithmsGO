package test2850

import (
	"math"
	"math/bits"
)

/**
 * @Description https://leetcode.cn/problems/minimum-moves-to-spread-stones-over-grid
 * idea: 状压dp + 全排列
 **/

func minimumMoves(g [][]int) int {
	var n = 3
	var ones = make([][]int, 0)
	var zeros = make([][]int, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if g[i][j] > 1 {
				for k := 2; k <= g[i][j]; k++ {
					ones = append(ones, []int{i, j})
				}
			} else if g[i][j] == 0 {
				zeros = append(zeros, []int{i, j})
			}
		}
	}
	var cal = func(i, j int) int {
		return abs(zeros[i][0]-ones[j][0]) + abs(zeros[i][1]-ones[j][1])
	}
	var m = len(ones)
	var l = 1 << m
	var dp = make([]int, l)
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt32
	}
	for i := 0; i < m; i++ {
		for s := l - 1; s >= 1; s-- { // 全排列
			if bits.OnesCount(uint(s)) != i+1 {
				continue
			}
			for j := 0; j < m; j++ { // 当前任一位置j与位置i匹配
				if s&(1<<j) != 0 {
					dp[s] = min(dp[s], dp[s^(1<<j)]+cal(i, j))
				}
			}
		}
	}
	return dp[l-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
