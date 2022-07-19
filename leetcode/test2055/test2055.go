package test2055

func platesBetweenCandles(s string, queries [][]int) []int {
	n, sum := len(s), 0
	var sums = make([]int, n)
	var left = make([]int, n)
	var right = make([]int, n)
	l, r := -1, -1
	for i := 0; i < n; i++ {
		if s[i] == '*' {
			sum++
		} else {
			l = i
		}
		sums[i] = sum
		left[i] = l
	}
	for i := n - 1; i >= 0; i-- {
		if s[i] == '|' {
			r = i
		}
		right[i] = r
	}
	var res = make([]int, len(queries))
	for i, q := range queries {
		x, y := right[q[0]], left[q[1]]
		if x >= 0 && y >= 0 && x < y {
			res[i] = sums[y] - sums[x]
		}
	}
	return res
}
