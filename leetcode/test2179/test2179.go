package test2179

/**
 * @Description https://leetcode.cn/problems/count-good-triplets-in-an-array/
 * idea: 三元组，常用方法： 枚举元组中的1-2个数，统计另外的数求最终解，
		常用到的操作为【单点】更新和【区间】查询，可以使用线段树/树状数组等
		具体到本题，求解在nums1和nums2中同样元素相对顺序相同的三元组, 考虑【固定】一个数组位置
		1.先存储nums1中每个数对应的索引
		2.再从左到右遍历nums2数组，根据1中的位置，使用树状数组累加对应的索引位置, 对于每个元素，
		  因为树状数组已经存储的元素在nums2中一定是当前元素左边的，而更新的实际位置又是nums1的位置
		  所以使用数状树状区间查询，当前元素在nums1中位置左边的元素，即为满足nums1和nums2同样元素相对顺序相等的元素个数
        3.从右到左遍历同理，根据乘法原理算出以当前元素为中心的三元组个数，计入答案即可
 **/

func goodTriplets(nums1 []int, nums2 []int) int64 {
	var n = len(nums1)
	var idx1 = make([]int, n)
	for i, v := range nums1 {
		idx1[v] = i + 1
	}
	var tree = NewTreeArr(n + 1)
	var left = make([]int, n)
	for i, v := range nums2 {
		var ti = idx1[v]
		left[i] = tree.QueryRange(1, ti)
		tree.UpdateRange(ti, ti, 1)
	}
	tree = NewTreeArr(n + 1)
	var res int64 = 0
	for i := n - 1; i >= 1; i-- {
		var ti = idx1[nums2[i]]
		res += int64(left[i] * tree.QueryRange(ti, n))
		tree.UpdateRange(ti, ti, 1)
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
