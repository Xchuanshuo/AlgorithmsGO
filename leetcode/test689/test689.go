package test689


func maxSumOfThreeSubarrays(nums []int, k int) []int {
	max, n := 0, len(nums)
	sum := make([]int, n+1)
	for i := 0;i < n;i++ {
		sum[i+1] = sum[i] + nums[i]
	}
	dp := make([][4]int, n+1)
	for j := 1;j <= 3;j++ {
		for i := j*k;i <= n;i++ {
			cur := sum[i] - sum[i-k]
			dp[i][j] = maxI(dp[i-1][j], dp[i-k][j-1] + cur)
			if j == 3 {
				max = maxI(max, dp[i][j])
			}
		}
	}
	res := make([]int, 3)
	for j := 3;j >= 1;j--{
		for i := 1;i <= n;i++ {
			if dp[i][j] == max {
				res[j-1] = i - k
				max -= sum[i] - sum[i-k]
				break
			}
		}
	}
	return res
}

func maxI(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//func maxSumOfThreeSubarrays(nums []int, k int) []int {
//	cur, n := 0, len(nums)
//	sum := make([]int, n+1)
//	for i := 0;i < n;i++ {
//		sum[i+1] = sum[i] + nums[i]
//	}
//	parent := make([][4]int, n+1)
//	dp := make([][4]int, n+1)
//	for j := 1;j <= 3;j++ {
//		for i := k;i <= n;i++ {
//			if i - k >= 0 {
//				cur = sum[i] - sum[i-k]
//			} else {
//				cur = 0
//			}
//			for p := 0;p <= i-k;p++ {
//				if dp[i][j] < dp[p][j-1] + cur {
//					dp[i][j] = dp[p][j-1] + cur
//					parent[i][j] = p
//				}
//			}
//		}
//	}
//	mIdx := 0
//	for i := 1;i <= n;i++ {
//		if dp[i][3] > dp[mIdx][3] {
//			mIdx = i
//		}
//	}
//	res := make([]int, 3)
//	res[2] = mIdx - k
//	res[1] = parent[mIdx][3] - k
//	res[0] = parent[parent[mIdx][3]][2] - k
//	return res
//}