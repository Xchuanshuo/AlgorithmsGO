package lcw370

func findChampion2(n int, edges [][]int) int {
	var degree = make([]int, n)
	for _, e := range edges {
		degree[e[1]]++
	}
	var res = -1
	var cnt0 = 0
	for i := 0; i < n; i++ {
		if degree[i] == 0 {
			cnt0++
			res = i
		}
	}
	if cnt0 > 1 {
		return -1
	}
	return res
}
