package lcw373

import "sort"

// 按limit距离分组, 每组单独从小到大
func lexicographicallySmallestArray(nums []int, limit int) []int {
	var as = make([][]int, 0)
	for i, v := range nums {
		as = append(as, []int{v, i})
	}
	sort.Slice(as, func(i, j int) bool {
		return as[i][0] < as[j][0]
	})
	var n = len(nums)
	var gs = make([][]int, 0)
	var ng = make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		ng = append(ng, as[i][1])
		if i == 0 || as[i][0]-as[i-1][0] > limit {
			gs = append(gs, ng)
			ng = make([]int, 0)
		}
	}
	var res = make([]int, n)
	for _, g := range gs {
		var vs = make([]int, 0)
		for _, idx := range g {
			vs = append(vs, nums[idx])
		}
		sort.Slice(g, func(i, j int) bool {
			return g[i] < g[j]
		})
		sort.Ints(vs)
		for j := len(g) - 1; j >= 0; j-- {
			res[g[j]] = vs[j]
		}
	}
	return res
}
