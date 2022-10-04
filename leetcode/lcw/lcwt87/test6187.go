package main

/**
 * @Description https://leetcode.cn/problems/minimum-money-required-before-transactions/
 * idea: 贪心, 假设除当前项目外 所有亏损已经计算完, 亏损+当前投资花费 即是最少值
 *	     依次将每个项目作为最后投资的项目 比较即可
 **/

func minimumMoney(transactions [][]int) int64 {
	var total int64 = 0
	for _, t := range transactions {
		if t[0] > t[1] {
			total += int64(t[0] - t[1])
		}
	}
	var res int64 = 0
	for _, t := range transactions {
		if t[0] > t[1] {
			total -= int64(t[0] - t[1])
		}
		res = maxInt64(res, total+int64(t[0]))
		if t[0] > t[1] {
			total += int64(t[0] - t[1])
		}
	}
	return res
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
