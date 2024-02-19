package lcwt118

/**
 * @Description https://leetcode.cn/contest/biweekly-contest-118/problems/find-maximum-non-decreasing-array-length/
 * idea: 单调队列优化dp
		 1.dp[i]表示使数组[0,i]变为非递减数组的最大长度
		   last[i]表示使数组[0,i]变为非递减数组情况下的最小和
		 2. dp[i] = dp[j] + 1, 且 sum[i]-sum[j] >= last[j]，即
			sum[i] >= last[j] + sum[j]，而dp[0..i]实际也满足非递减, 总是能将一个元素合并到之前的和上
			所以若存在 k < j 且 last[k] + sum[k] <= last[j] + sum[j]，则last[k] + sum[k]后续无需使用
		 3.使用单调递增队列维护即可，队首越大，则last[j]越小
 **/

func findMaximumLength(nums []int) int {
	var n = len(nums)
	var sum = make([]int, n+1)
	var dp = make([]int, n+1)
	var last = make([]int, n+1)
	var q = []int{0}
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + nums[i-1]
		// 对于求答案，和越大能保证当前结尾越小， 剔除队首较小和
		for len(q) > 1 && last[q[1]]+sum[q[1]] <= sum[i] {
			q = q[1:]
		}
		last[i] = sum[i] - sum[q[0]]
		dp[i] = dp[q[0]] + 1
		var cur = last[i] + sum[i]
		// 对于入队尾，和较大的情况下不会比后面较小的更优，剔除队尾较大和
		for len(q) > 0 && last[q[len(q)-1]]+sum[q[len(q)-1]] >= cur {
			q = q[0 : len(q)-1]
		}
		q = append(q, i)
	}
	return dp[n]
}
