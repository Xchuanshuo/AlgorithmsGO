package lcw330

/**
 * @Description https://leetcode.cn/problems/count-increasing-quadruplets/
// 四元组，i < j < k < l, nums[i] < nums[k] < nums[j] < nums[l]
//  枚举索引j,k, 树状数组求 1.j左边小于k的元素个数 2.k右边大于j的元素个数
 **/
// greater[x][k]
func countQuadruplets(nums []int) int64 {
	var n = len(nums)
	var a1 = NewTreeArr(n + 1)
	var res int64 = 0
	for j := 0; j < n; j++ {
		var a2 = NewTreeArr(n + 1)
		for k := n - 1; k >= 0; k-- {
			if j == k {
				break
			}
			a2.UpdateRange(nums[k], nums[k], 1)
			if nums[j] < nums[k] {
				continue
			}
			var cnt1 = a1.query(nums[k] - 1)
			var cnt2 = a2.QueryRange(nums[j]+1, n)
			res += int64(cnt1 * cnt2)
		}
		a1.UpdateRange(nums[j], nums[j], 1)
	}
	return res
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
