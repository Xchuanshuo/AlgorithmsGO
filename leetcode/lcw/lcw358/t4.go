package lcw358

import (
	"math"
	"sort"
)

// 分解质因数 + 排序 + 单调栈 + 前缀和 + 快速幂,
// 计算每个元素覆盖的区间数, 为左右两边第一个大于其分数的元素位置

var cnt = make(map[int]int)

func init() {
	for i := 1; i <= int(1e5); i++ {
		cnt[i] = f(i)
	}
}

var M = int(1e9) + 7

func maximumScore(nums []int, k int) int {
	var n = len(nums)
	var s = [][]int{{n, math.MaxInt32}}
	var right = make([]int, n)
	for i := n - 1; i >= 0; i-- {
		var cur = cnt[nums[i]]
		for s[len(s)-1][1] <= cur {
			s = s[0 : len(s)-1]
		}
		right[i] = max(1, s[len(s)-1][0]-i)
		s = append(s, []int{i, cur})
	}
	s = [][]int{{-1, math.MaxInt32}}
	var a = make([][]int, 0)
	for i := 0; i < n; i++ {
		var cur = cnt[nums[i]]
		for s[len(s)-1][1] < cur {
			s = s[0 : len(s)-1]
		}
		a = append(a, []int{nums[i], right[i] * max(i-s[len(s)-1][0], 1)})
		s = append(s, []int{i, cur})
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i][0] < a[j][0]
	})
	var j = len(a) - 1
	var res = 1
	for k > 0 && j >= 0 {
		var x = a[j][0]
		if k >= a[j][1] {
			res = (res*fastPow(x, a[j][1]) + M) % M
			k -= a[j][1]
			j--
		} else {
			res = (res*fastPow(x, k) + M) % M
			k = 0
		}
	}
	return res
}

func fastPow(a, k int) int {
	if k == 0 {
		return 1
	}
	a %= M
	if k%2 == 0 {
		val := fastPow(a, k/2)
		return (val * val) % M
	} else {
		return a * fastPow(a, k-1) % M
	}
}

var f = func(n int) int {
	var facs = make(map[int]bool)
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			for n%i == 0 {
				n /= i
				facs[i] = true
			}
		}
	}
	if n > 1 {
		facs[n] = true
	}
	return len(facs)
}
