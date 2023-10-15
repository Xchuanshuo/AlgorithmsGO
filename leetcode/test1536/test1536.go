package test1536

/**
 * @Description https://leetcode.cn/problems/minimum-swaps-to-arrange-a-binary-grid
 * idea: 1.预处理，计算每行结尾连续0的个数 2.模拟冒泡排序，查找满足当前行要求的最近位置
 **/

func minSwaps(g [][]int) int {
	var n = len(g)
	var a = make([]int, 0)
	for i := 0; i < n; i++ {
		var cnt = 0
		for j := n - 1; j >= 0; j-- {
			if g[i][j] != 0 {
				break
			}
			cnt++
		}
		a = append(a, cnt)
	}
	var res = 0
	for i := 0; i < n; i++ {
		var idx = -1
		for j := i; j < n; j++ {
			if a[j] >= n-i-1 {
				idx = j
				break
			}
		}
		if idx == -1 {
			return -1
		}
		for j := idx; j > i; j-- {
			a[j], a[j-1] = a[j-1], a[j]
			res++
		}
	}
	return res
}
