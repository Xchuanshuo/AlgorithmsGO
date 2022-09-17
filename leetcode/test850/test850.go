package test850

import "sort"

/**
 * @Description
 * idea: 扫描线 1.将x点离散 从小到大排序 2.找到覆盖x区间的所有矩形 计算总高度
 *			                           可使用线段树快速判断区间是否被覆盖
 **/

const MOD = int(1e9 + 7)

func rectangleArea(rectangles [][]int) int {
	var xs = make([]int, 0)
	for _, rec := range rectangles {
		xs = append(xs, rec[0], rec[2])
	}
	sort.Slice(xs, func(i, j int) bool {
		return xs[i] < xs[j]
	})
	var res = 0
	for i := 1; i < len(xs); i++ {
		x1, x2 := xs[i-1], xs[i]
		var hs = make([][]int, 0)
		for _, rec := range rectangles {
			if rec[0] <= x1 && rec[2] >= x2 {
				hs = append(hs, []int{rec[1], rec[3]})
			}
		}
		sort.Slice(hs, func(i, j int) bool {
			if hs[i][0] != hs[j][0] {
				return hs[i][0] < hs[j][0]
			}
			return hs[i][1] < hs[j][1]
		})
		total, down, up := 0, -1, -1
		for _, h := range hs {
			if h[0] > up {
				total += up - down
				down, up = h[0], h[1]
			} else if h[1] > up {
				up = h[1]
			}
		}
		total += up - down
		res = (res + total*(xs[i]-xs[i-1])) % MOD
	}
	return res
}
