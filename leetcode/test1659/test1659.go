package main

import (
	"fmt"
	"math"
)

/**
 * @Description https://leetcode.cn/problems/maximize-grid-happiness/
 * idea: 状压DP + 记忆化搜索
		 由于每个网格可以有三种安排情况，不安排任何人、安排内向者、安排外向者，所以考虑用【3进制】进行编码，
		 记录格子的状态，枚举所有不同的网格状态计算结果即可。

		 首先考虑将m*n网格打平后记录所有格子的状态, 此时总状态数为3^(m*n)=1e11，时间复杂度过高；
         继续考虑，对于每个格子，实际只关心与它相邻格子的状态。从第0行开始处理，对于每个格子，只需要关心它左边的格子和
		 上边的格子的安排方式。所以考虑枚举每行的状态，总数为3^n，由于不需要每个人都安排在网格中，所以可以将当前的行当作
  		 最终处理的行，因此提前预处理所有【当前状态行的得分】与【相邻行之间状态的贡献分】，累加后即为当前行最终得分，接着继续
		 处理下一行，过程中取最大贡献即可。

		 整体时间复杂度为O((m*maxI*maxE+n)*3^(2n))，其中记忆化搜索后使用dfs处理总的状态个数为 m*maxI*maxE*3^n，
		 每个状态计算需要花费3^n，最终时间为m*maxI*maxE*3^(2n)；预处理不同状态行之间的时间复杂度为 3^(2*n)*n。
 **/

// 记录不同相邻状态时贡献的分数 0-不安排 1-内向 2-外向
var score = [][]int{{0, 0, 0}, {0, -60, -10}, {0, -10, 40}}

func getMaxGridHappiness(m int, n int, maxI int, maxE int) int {
	// total表示每列的状态总数
	var total = int(math.Pow(3.0, float64(n)))
	var rowV = make([][]int, total)
	var colV = make([]int, total)
	// 记录不同状态下内向和外向的人数
	var cntI = make([]int, total)
	var cntE = make([]int, total)
	for i := 0; i < total; i++ {
		rowV[i] = make([]int, total)
	}
	// 1.预处理不同状态行的内向、外向人数和对应的分数
	for i := 0; i < total; i++ {
		var mask = i
		var pre = 0
		for j := 0; j < n; j++ {
			var cur = mask % 3
			mask /= 3
			if cur == 1 {
				cntI[i]++
				colV[i] += 120
			} else if cur == 2 {
				cntE[i]++
				colV[i] += 40
			}
			// 累加相邻列之间的贡献
			colV[i] += score[cur][pre]
			pre = cur
		}
	}
	// 2.预处理所有相邻行之间的贡献分数
	for i := 0; i < total; i++ {
		for j := 0; j < total; j++ {
			maskI, maskJ := i, j
			for k := 0; k < n; k++ {
				p1, p2 := maskI%3, maskJ%3
				maskI, maskJ = maskI/3, maskJ/3
				rowV[i][j] += score[p1][p2]
			}
		}
	}
	var mem = getInitMem(m, total, maxI, maxE)
	// 子问题为从第i行开始，前一个状态为pre，剩余内向ri人，外向re人的情况下最大的幸福感是多少
	var dfs func(i, pre, ri, re int) int
	dfs = func(i, pre, ri, re int) int {
		if i == m || (ri == 0 && re == 0) {
			return 0
		}
		if mem[i][pre][ri][re] != -1 {
			return mem[i][pre][ri][re]
		}
		var res = 0
		for cur := 0; cur < total; cur++ {
			if ri >= cntI[cur] && re >= cntE[cur] {
				res = max(res, colV[cur]+rowV[pre][cur]+dfs(i+1, cur, ri-cntI[cur], re-cntE[cur]))
			}
		}
		mem[i][pre][ri][re] = res
		return res
	}
	return dfs(0, 0, maxI, maxE)
}

func getInitMem(m int, total int, maxI, maxE int) [][][][]int {
	var mem = make([][][][]int, m)
	for i := 0; i < len(mem); i++ {
		mem[i] = make([][][]int, total)
		for j := 0; j < len(mem[i]); j++ {
			mem[i][j] = make([][]int, maxI+1)
			for k := 0; k < len(mem[i][j]); k++ {
				mem[i][j][k] = make([]int, maxE+1)
				for x := 0; x < len(mem[i][j][k]); x++ {
					mem[i][j][k][x] = -1
				}
			}
		}
	}
	return mem
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var res = getMaxGridHappiness(5, 5, 3, 6)
	fmt.Println(res)
}
