package test769

func maxChunksToSorted(arr []int) int {
	m, res := 0, 0
	for i := 0; i < len(arr); i++ {
		m = max(m, arr[i])
		if i == m {
			res++
		}
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
