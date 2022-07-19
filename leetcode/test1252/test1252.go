package test1252

func oddCells(m int, n int, indices [][]int) int {
	var row = make([]int, m)
	var col = make([]int, n)
	for _, idx := range indices {
		row[idx[0]]++
		col[idx[1]]++
	}
	var res = 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (row[i]+col[j])&1 == 1 {
				res++
			}
		}
	}
	return res
}
