package test2022

func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m * n {
		return make([][]int, 0)
	}
	idx := 0
	res := make([][]int, m)
	for i := 0;i < m;i++ {
		res[i] = make([]int, n)
		for j := 0;j < n;j++ {
			res[i][j] = original[idx]
			idx++
		}
	}
	return res
}