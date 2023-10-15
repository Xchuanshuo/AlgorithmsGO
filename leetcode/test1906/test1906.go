package test1906

/**
 * @Description https://leetcode.cn/problems/minimum-absolute-difference-queries/
 * idea: 前缀和 + 枚举值域
 **/

func minDifference(nums []int, queries [][]int) []int {
	var n = len(nums)
	var sum = make([][101]int, n+1)
	for i := 0; i < n; i++ {
		for j := 1; j <= 100; j++ {
			sum[i+1][j] = sum[i][j]
		}
		sum[i+1][nums[i]] += 1
	}
	var res = make([]int, len(queries))
	for i, q := range queries {
		var cur = 1000
		var pre = -1
		for j := 1; j <= 100; j++ {
			var cnt = sum[q[1]+1][j] - sum[q[0]][j]
			if cnt == 0 {
				continue
			}
			if pre != -1 {
				cur = min(cur, j-pre)
			}
			pre = j
		}
		if cur == 1000 {
			res[i] = -1
		} else {
			res[i] = cur
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
