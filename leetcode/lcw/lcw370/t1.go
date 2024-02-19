package lcw370

func findChampion(g [][]int) int {
	var n = len(g)
	var res = 0
	for i := 0; i < n; i++ {
		var cnt = 0
		for j := 0; j < n; j++ {
			if g[j][i] != 1 {
				cnt++
			}
		}
		if cnt == n {
			res = i
			break
		}
	}
	return res
}
