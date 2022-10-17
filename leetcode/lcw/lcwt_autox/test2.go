package lwc_autox

func honeyQuotes(handle [][]int) []float64 {
	var res = make([]float64, 0)
	sum1, sum2 := 0.0, 0.0
	var cnt = 0
	for _, h := range handle {
		if h[0] == 1 {
			sum1 += float64(h[1])
			sum2 += float64(h[1] * h[1])
			cnt++
		} else if h[0] == 2 {
			sum1 -= float64(h[1])
			sum2 -= float64(h[1] * h[1])
			cnt--
		} else if h[0] == 3 {
			if cnt == 0 {
				res = append(res, -1)
			} else {
				res = append(res, sum1/float64(cnt))
			}
		} else if h[0] == 4 {
			if cnt == 0 {
				res = append(res, -1)
			} else {
				res = append(res, (sum2-sum1*sum1/float64(cnt))/float64(cnt))
			}
		}
	}
	return res
}
