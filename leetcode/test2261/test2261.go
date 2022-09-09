package test2261

/**
 * @Description https://leetcode.cn/problems/k-divisible-elements-subarrays/submissions/
 * idea: 字符串hash, 为避免重复，计算两个hash值拼接即可
 **/

func countDistinct(nums []int, k int, p int) int {
	var n = len(nums)
	var set = make(map[int64]bool)
	m, p1, p2 := int(1e9)+7, 37, 41
	for i := 0; i < n; i++ {
		var cnt = 0
		h1, h2 := 0, 0
		for j := i; j < n; j++ {
			if nums[j]%p == 0 {
				cnt++
			}
			if cnt > k {
				break
			}
			h1 = (13*p1*h1 + nums[j]) % m
			h2 = (13*p2*h2 + nums[j]) % m
			var h = int64(h1<<32) | int64(h2)
			set[h] = true
		}
	}
	return len(set)
}
