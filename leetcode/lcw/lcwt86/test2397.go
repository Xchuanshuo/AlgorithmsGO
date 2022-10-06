package lcwt86

import "math/bits"

func maximumRows(matrix [][]int, cols int) int {
	m, n := len(matrix), len(matrix[0])
	var arr = make([]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				arr[i] |= 1 << j
			}
		}
	}
	var length = 1 << cols
	var res = 0
	for i := 1; i < length; i++ {
		var cnt = bits.OnesCount(uint(i))
		if cnt != cols {
			continue
		}
		var cur = 0
		for _, a := range arr {
			if i&a == i {
				cur++
			}
		}
		if cur > res {
			res = cur
		}
	}
	return res
}
