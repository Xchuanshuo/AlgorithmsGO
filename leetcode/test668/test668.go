package test668

func findKthNumber(m int, n int, k int) int {
	l, r := 1, m*n
	for l <= r {
		mid := l + (r-l)/2
		cnt := getCnt(m, n, mid)
		if cnt < k {
			l = mid + 1
		} else {
			if mid == l || getCnt(m, n, mid-1) < k {
				return mid
			}
			r = mid - 1
		}
	}
	return -1
}

func getCnt(m, n int, mid int) int {
	var cnt = 0
	for i := 1; i <= m && i <= mid; i++ {
		cnt += min(n, mid/i)
	}
	return cnt
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
