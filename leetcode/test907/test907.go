package test907

/**
 * @Description https://leetcode.cn/problems/sum-of-subarray-minimums/
 * idea: 计算每个数对结果的贡献 -> 找到两个比它小的端点
 **/

func sumSubarrayMins(arr []int) int {
	var M = int(1e9) + 7
	arr = append(arr, 0)
	var s = []int{-1}
	n, res := len(arr), 0
	for i := 0; i < n; i++ {
		for len(s) > 1 && arr[s[len(s)-1]] > arr[i] {
			var ti = s[len(s)-1]
			s = s[0 : len(s)-1]
			res = (res + (i-ti)*(ti-s[len(s)-1])*arr[ti]) % M
		}
		s = append(s, i)
	}
	return res
}
