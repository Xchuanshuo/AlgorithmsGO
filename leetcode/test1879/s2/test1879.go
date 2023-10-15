package s2

import (
	"math"
	"math/bits"
)

/**
 * @Description https://leetcode.cn/problems/minimum-xor-sum-of-two-arrays/
 * idea: 状态压缩枚举排列，dp[i][s]表示nums1的前i个数和nums2选中状态为s时的最小异或和

 **/

func minimumXORSum(nums1 []int, nums2 []int) int {
	var n = len(nums1)
	var l = 1 << n
	var dp = make([]int, l)
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt32
	}
	for i := 0; i < n; i++ {
		for s := l - 1; s >= 0; s-- {
			if bits.OnesCount(uint(s)) != i+1 {
				continue
			}
			for j := 0; j < n; j++ {
				if s&(1<<j) == 0 {
					continue
				}
				dp[s] = min(dp[s], dp[s^(1<<j)]+(nums1[i]^nums2[j]))
			}
		}
	}
	return dp[l-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
