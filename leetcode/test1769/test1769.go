package test1769

/**
 * @Description https://leetcode.cn/problems/minimum-number-of-operations-to-move-all-balls-to-each-box
 * idea: 前缀和
 **/

func minOperations(boxes string) []int {
	var n = len(boxes)
	var sum = make([]int, n+1)
	var cnt = make([]int, n+1)
	for i := 0; i < n; i++ {
		var v = int(boxes[i] - '0')
		cnt[i+1] = cnt[i] + v
		sum[i+1] = sum[i] + v*(i+1)
	}
	var res = make([]int, n)
	for i := 0; i < n; i++ {
		var left = cnt[i]*(i+1) - sum[i]
		var right = (sum[n] - sum[i+1]) - (cnt[n]-cnt[i+1])*(i+1)
		res[i] = left + right
	}
	return res
}
