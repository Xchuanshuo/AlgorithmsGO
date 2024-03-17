package lcw386

func largestSquareArea(bl [][]int, rt [][]int) int64 {
	var recs = make([]int, 0)
	var n = len(bl)
	for i := 0; i < n; i++ {
		recs = append(recs, i)
	}
	var calW = func(i, j int) int {
		bl1, rt1 := bl[recs[i]], rt[recs[i]]
		bl2, rt2 := bl[recs[j]], rt[recs[j]]
		// 矩形相交判定, 左侧x坐标取较大的，右侧x坐标取较小的；y坐标上下侧同理
		x0, y0 := max(bl1[0], bl2[0]), max(bl1[1], bl2[1])
		x1, y1 := min(rt1[0], rt2[0]), min(rt1[1], rt2[1])
		var w = min(x1-x0, y1-y0)
		if w < 0 {
			return 0
		}
		return w
	}
	var res = 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			res = max(res, calW(i, j))
		}
	}
	return int64(res * res)
}
