package test2333

/**
 * @Description https://leetcode.cn/problems/minimum-sum-of-squared-difference/
 * idea:
		1.由于对于nums1和nums2的修改可以+1,-1，所以nums1和nums2可以等价转换，即总修改次数为k1+k2
		2.要使总平方差值最小，优先考虑减小较大的abs(nums1[i]-nums2[i])，所以对绝对值进行排序
		3.前缀和预先计算前面所有位置的平方差值和。
	    4.倒序遍历，对于后面每一个值，考虑转换为当前值，过程中计算min即可
 **/

import "sort"

func minSumSquareDiff(nums1 []int, nums2 []int, k1 int, k2 int) int64 {
	var n = len(nums2) + 1
	var a = []int{0}
	for i := 0; i < n-1; i++ {
		a = append(a, abs(nums1[i]-nums2[i]))
	}
	sort.Ints(a)
	var sum = make([]int64, n+1)
	for i, v := range a {
		sum[i+1] = sum[i] + int64(v*v)
	}
	var res = sum[n]
	var total = k1 + k2
	var r = a[n-1]
	for i := n - 2; i >= 0; i-- {
		var cnt = n - i - 1
		var cost = cnt * (r - a[i])
		if total < cost {
			var sub = total / cnt
			var ext = total % cnt
			res = min(res, int64((cnt-ext)*(r-sub)*(r-sub)+ext*(r-sub-1)*(r-sub-1))+sum[i+1])
			break
		}
		res = min(res, sum[i+1]+int64(a[i]*a[i]*cnt))
		r = a[i]
		total -= cost
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
