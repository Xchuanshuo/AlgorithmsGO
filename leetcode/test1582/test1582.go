package test1582

func numSpecial(mat [][]int) int {
	m, n := len(mat), len(mat[0])
	var rows = make([]int, m)
	var cols = make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 {
				rows[i]++
				cols[j]++
			}
		}
	}
	var res = 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 && rows[i] == 1 && cols[j] == 1 {
				res++
			}
		}
	}
	return res
}
