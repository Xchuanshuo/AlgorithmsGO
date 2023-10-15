package lcwt113

func countPairs(coordinates [][]int, k int) int {
	var mp = make(map[T]int)
	var res = 0
	for _, c := range coordinates {
		x, y := c[0], c[1]
		for i := 0; i <= k; i++ {
			tx, ty := x^i, y^(k-i)
			res += mp[T{tx, ty}]
		}
		mp[T{x, y}]++
	}
	return res
}

type T struct {
	x, y int
}

//输入：
//[[27,94],[61,68],[47,0],[100,4],[127,89],[61,103],[26,4],[51,54],[91,26],[98,23],[80,74],[19,93]]
//95
//输出：
//4
//预期：
//5
