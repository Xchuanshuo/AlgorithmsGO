package test1774

/**
* @Description https://leetcode.cn/problems/closest-dessert-cost/
* idea: 动态规划(0-1背包), dp[j]表示能否凑到背包大小为j容量
        1.大于target则记录最小的取法
        2.小于等于target则看状态能否由【不加配料】后转移过来
		3.初始条件: 基料大小的背包能直接凑齐
		时间复杂度: O(target*len(t))
**/

func closestCost1(base []int, t []int, target int) int {
	t = append(t, t...)
	_, n := len(base), len(t)
	var dp = make([]bool, target+1)
	var res = int(1e9)
	for _, b := range base {
		if b > target {
			res = min(res, b)
		} else {
			dp[b] = true
		}
	}
	for i := 0; i < n; i++ {
		for j := target; j >= 0; j-- {
			if j+t[i] > target && dp[j] {
				res = min(res, j+t[i])
			}
			if j+t[i] <= target {
				dp[j+t[i]] = dp[j] || dp[j+t[i]]
			}
		}
	}
	for i := target; i >= 0; i-- {
		if dp[i] && target-i <= res-target {
			return i
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
