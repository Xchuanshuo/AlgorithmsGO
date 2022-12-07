package test1775

/**
 * @Description https://leetcode.cn/problems/equal-sum-arrays-with-minimum-number-of-operations/
 * idea: 贪心, 1.计算两个数组的差值 2.统计两个数组增量的个数 3.依次用最大增量依次抵消差值 即可计算最少次数
 **/

func minOperations(nums1 []int, nums2 []int) int {
	n1, n2 := len(nums1), len(nums2)
	if n1 > 6*n2 || n2 > 6*n1 {
		return -1
	}
	var d = 0
	for _, v := range nums1 {
		d += v
	}
	for _, v := range nums2 {
		d -= v
	}
	if d < 0 {
		d = -d
		nums1, nums2 = nums2, nums1
	}
	var cnt = make([]int, 6)
	for _, v := range nums1 {
		cnt[v-1]++
	}
	for _, v := range nums2 {
		cnt[6-v]++
	}
	var res = 0
	for i := 5; i >= 0; i-- {
		if i*cnt[i] >= d {
			res += (d + i - 1) / i
			break
		}
		res += cnt[i]
		d -= i * cnt[i]
	}
	return res
}
