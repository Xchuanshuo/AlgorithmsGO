package test1184

func distanceBetweenBusStops(distance []int, s int, d int) int {
	var n = len(distance)
	var sums = make([]int, n+1)
	for i := 0; i < n; i++ {
		sums[i+1] = sums[i] + distance[i]
	}
	if s > d {
		s, d = d, s
	}
	return min(sums[d]-sums[s], sums[n]-(sums[d]-sums[s]))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
