package test747

func dominantIndex(nums []int) int {
	res, m1, m2 := 0, 0, 0
	for idx, v := range nums {
		if v > m1 {
			res = idx
			m2 = m1
			m1 = v
		} else if v > m2 {
			m2 = v
		}
	}
	if m1 >= 2*m2 { return res }
	return -1
}