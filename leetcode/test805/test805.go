package test805

/**
 * @Description https://leetcode.cn/problems/split-array-with-same-average/
 * idea:
 *      背包问题 能否均值分割 -> 能否 凑满和为j的背包 使得 j/k = sum/n
 *      对于和j 可能有多个k, 用状态压缩, 1在第k位表示可以使用k个数字凑成和j
 *      只需满足 j*n%sum = 0, 且k能从 j-nums[i]的和转移过来
 *		初始条件: 使用1个数凑满和为0的背包
 **/

func splitArraySameAverage(nums []int) bool {
	n, sum := len(nums), 0
	if n == 1 {
		return false
	}
	for _, num := range nums {
		sum += num
	}
	if sum == 0 {
		return true
	}
	var total = sum / 2
	var dp = make([]int, total+1)
	dp[0] = 1
	for i := 0; i < n; i++ {
		for j := total; j >= nums[i]; j-- {
			dp[j] |= dp[j-nums[i]] << 1
			if j*n%sum == 0 && (1<<(j*n/sum)&dp[j]) > 0 {
				if float64(sum-j)/float64(n-(j*n/sum)) == float64(j)/float64(j*n/sum) {
					return true
				}
			}
		}
	}
	return false
}
