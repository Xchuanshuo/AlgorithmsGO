package s1

/**
 * @Description https://leetcode.cn/problems/minimum-number-of-operations-to-make-all-array-elements-equal-to-1/
 * idea: 1.原数组存在cnt个11，直接构造，答案为n-cnt
		 2.不存在1，枚举所有区间，若某个区间计算出最大公约数1，
		   直接使用这个1进行构造, 计算最小值即可
		   时间复杂度为O(n^2)
 **/

func minOperations(nums []int) int {
	var n = len(nums)
	var cnt = 0
	for _, v := range nums {
		if v == 1 {
			cnt++
		}
	}
	if cnt > 0 {
		return n - cnt
	}
	var res = 2 * n
	for i := 0; i < n; i++ {
		var cur = nums[i]
		for j := i + 1; j < n; j++ {
			cur = gcd(cur, nums[j])
			if cur == 1 {
				res = min(res, j-i+n-1)
			}
		}
	}
	if res == 2*n {
		return -1
	}
	return res
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
