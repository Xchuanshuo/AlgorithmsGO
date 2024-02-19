package lcwt118

/**
 * @Description https://leetcode.cn/problems/minimum-number-of-coins-for-fruits/
 * idea: 单调队列优化版本
 **/

func minimumCoins1(prices []int) int {
	var n = len(prices)
	var q = make([][]int, 0)
	q = append(q, []int{n + 1, 0})
	for i := n; i >= 1; i-- {
		for len(q) > 0 && q[0][0] > 2*i+1 {
			q = q[1:]
		}
		// 从位置i，购买完所有水果的最小开销
		var cur = prices[i-1] + q[0][1]
		for len(q) > 0 && q[len(q)-1][1] >= cur {
			q = q[0 : len(q)-1]
		}
		q = append(q, []int{i, cur})
	}
	return q[len(q)-1][1]
}
