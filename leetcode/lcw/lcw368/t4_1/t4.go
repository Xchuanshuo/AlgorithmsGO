package main

import (
	"fmt"
	"math"
	time2 "time"
)

// 1.任意区间修改为半回文串的开销
// 2.dp[i][k]前i个数分成k段的最少修改次数

var facs = make([][]int, 201)

func init() {
	for i := 1; i <= 200; i++ {
		for j := 1; j < i; j++ {
			if i%j == 0 {
				facs[i] = append(facs[i], j)
			}
		}
	}
}

func minimumChanges(s string, K int) int {
	var n = len(s)
	var cal = func(l, r int) int {
		var maxD = r - l + 1
		var cost = maxD
		for _, d := range facs[maxD] {
			var gs = make([][]int, d)
			for i := 0; i < maxD; i++ {
				var rr = i % d
				gs[rr] = append(gs[rr], i+l)
			}
			var cnt = 0
			for _, g := range gs {
				var l = len(g)
				for i := 0; i < l/2; i++ {
					if s[g[i]] != s[g[l-i-1]] {
						cnt++
					}
				}
			}
			cost = min(cost, cnt)
		}
		return cost
	}
	var cost = make([][]int, n+1)
	for i := 0; i < len(cost); i++ {
		cost[i] = make([]int, n+1)
	}
	for l := 2; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			var j = i + l - 1
			cost[i+1][j+1] = cal(i, j)
		}
	}
	var dp = make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, K+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = math.MaxInt32
		}
		dp[i][1] = cost[1][i]
	}
	for k := 2; k <= K; k++ {
		for i := k * 2; i <= n; i++ {
			for j := i - 2; j >= 2; j-- {
				dp[i][k] = min(dp[i][k], dp[j][k-1]+cost[j+1][i])
			}
		}
	}
	return dp[n][K]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func timeMeasurement(start time2.Time) {
	elapsed := time2.Since(start)
	fmt.Printf("Execution time: %dms", elapsed.Milliseconds())
}

func main() {
	var s = "dcbadcbadcbadcbadcbadcbadcbadcbadcbadcbadcbadcbadcbadcbadcbadcbadcbadcbadcbadcbadcbadcbadcbadcbadcbadcbadcbadcbadcbadcbadcbadcbadcbadcbadcbadcbadcbadcbadcbadcbadcbadcbadcbadcbadcbadcbadcbadcbadcbadcba"
	var k = 100
	defer timeMeasurement(time2.Now())
	minimumChanges(s, k)
}
