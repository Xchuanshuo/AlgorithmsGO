package test2659

import (
	"sort"
)

/**
 * @Description https://leetcode.cn/problems/make-array-empty/
 * idea: 问题可转换为下标0开始，从循环数组中依次删除最小值，问删除完花费的总步数是多少
        因为删除是按元素大小顺序, 所以需要排序。而计算移动步数又依赖原顺序，所以构建一个
		索引数组排序即可。最终问题就变成求解两个相邻大小的元素，对应原数组中间有多少个元素被删除了。
		使用树状数组，元素删除时做标记即可。 设前一个元素为pre，当前为i，在原数组中有两种情况
		1.pre<=i, 此时移动步数为 i-pre-query(pre, i)
		2.pre>i, 此时移动 (n-pre-query(pre, n)) + (i-query(1, i))
 **/

func countOperationsToEmptyArray(nums []int) int64 {
	var n = len(nums)
	var ids = make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = i + 1
	}
	sort.Slice(ids, func(i, j int) bool {
		return nums[ids[i]-1] < nums[ids[j]-1]
	})
	var tree = NewTreeArr(n + 1)
	pre, res := 1, n
	for _, i := range ids {
		if pre <= i {
			res += i - pre - tree.QueryRange(pre, i)
		} else {
			res += (n - pre - tree.QueryRange(pre, n)) + (i - tree.QueryRange(1, i))
		}
		tree.UpdateRange(i, i, 1)
		pre = i
	}
	return int64(res)
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
