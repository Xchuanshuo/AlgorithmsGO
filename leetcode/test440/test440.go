package test440

func getCount(prefix int, n int) int {
	left, right := prefix, prefix+1
	var cnt = 0
	for left <= n {
		cnt += min(right, n+1) - left
		left *= 10
		right *= 10
	}
	return cnt
}

// https://leetcode-cn.com/problems/k-th-smallest-in-lexicographical-order/solution/ben-ti-shi-shang-zui-wan-zheng-ju-ti-de-shou-mo-sh/
func findKthNumber(n int, k int) int {
	p, prefix := 1, 1
	for p < k {
		var count = getCount(prefix, n)
		if p+count <= k {
			// 到右边查找
			prefix++
			p += count
		} else {
			// 到子树查找
			prefix *= 10
			p++
		}
	}
	return prefix
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
