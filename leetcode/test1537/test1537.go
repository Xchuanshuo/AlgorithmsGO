package test1537

var M = int64(1e9) + 7

func maxSum(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	var f = make([]int64, m+1)
	var g = make([]int64, n+1)
	i, j := 1, 1
	for i <= m || j <= n {
		if i <= m && j <= n {
			if nums1[i-1] < nums2[j-1] {
				f[i] = f[i-1] + int64(nums1[i-1])
				i++
			} else if nums1[i-1] > nums2[j-1] {
				g[j] = g[j-1] + int64(nums2[j-1])
				j++
			} else {
				var v = max(f[i-1], g[j-1]) + int64(nums1[i-1])
				f[i], g[j] = v, v
				i++
				j++
			}
		} else if i <= m {
			f[i] = f[i-1] + int64(nums1[i-1])
			i++
		} else {
			g[j] = g[j-1] + int64(nums2[j-1])
			j++
		}
	}
	return int(max(f[m], g[n]) % M)
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
