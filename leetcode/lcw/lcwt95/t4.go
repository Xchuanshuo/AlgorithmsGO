package lcwt95

/**
 * @Description https://leetcode.cn/contest/biweekly-contest-95/problems/maximize-the-minimum-powered-city/
 * idea: 二分 + 差分数组区间修改和查询
 **/

func maxPower(stations []int, rg int, k int) int64 {
	var n = len(stations)
	var srcTree = NewTreeArr(n + 1)
	for i, v := range stations {
		l, r := max(1, i+1-rg), min(n, i+1+rg)
		srcTree.UpdateRange(l, r, v)
	}
	var check = func(target int) bool {
		var tree = srcTree.Clone()
		var total = 0
		for i := 1; i <= n; i++ {
			var cur = tree.QueryRange(i, i)
			if cur < target {
				var dif = target - cur
				total += dif
				if total > k {
					return false
				}
				l, r := i, min(n, i+2*rg)
				tree.UpdateRange(l, r, dif)
			}
		}
		return true
	}

	l, r := 0, int(2e10)
	for l <= r {
		var mid = l + (r-l)/2
		if !check(mid) {
			r = mid - 1
		} else {
			if mid == int(2e10) || !check(mid+1) {
				return int64(mid)
			}
			l = mid + 1
		}
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type TreeArr struct {
	dif1, dif2 []int
}

func NewTreeArr(n int) *TreeArr {
	return &TreeArr{dif1: make([]int, n), dif2: make([]int, n)}
}
func (t *TreeArr) QueryRange(l, r int) int {
	return t.query(r) - t.query(l-1)
}
func (t *TreeArr) UpdateRange(l, r int, val int) {
	t.update(l, val)
	t.update(r+1, -val)
}
func (t *TreeArr) query(i int) int {
	var x = i
	var res = 0
	for i > 0 {
		res += (x+1)*t.dif1[i] - t.dif2[i]
		i -= t.lowbit(i)
	}
	return res
}
func (t *TreeArr) update(i int, val int) {
	x, n := i, len(t.dif1)
	for i < n {
		t.dif1[i] += val
		t.dif2[i] += x * val
		i += t.lowbit(i)
	}
}
func (t *TreeArr) lowbit(x int) int {
	return x & -x
}

func (t *TreeArr) Clone() *TreeArr {
	var cloneDif1 = make([]int, len(t.dif1))
	var cloneDif2 = make([]int, len(t.dif2))
	copy(cloneDif1, t.dif1)
	copy(cloneDif2, t.dif2)
	return &TreeArr{cloneDif1, cloneDif2}
}
