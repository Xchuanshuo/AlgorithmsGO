package lcw335

/**
 * @Description https://leetcode.cn/problems/split-the-array-to-make-coprime-products/
 * idea: 求两边乘积互质 -> 分割线两边 不存在相同的质因数 -> 求解每个数分解质因数后所能覆盖的最大区间
		 从左到右枚举所有索引位置，寻找分割点，若前面元素的质因数所覆盖的最大区间小于当前分割点，
         当前分割点即为答案，实际题目为求解跳跃游戏
 **/

import (
	"math"
)

func findValidSplit(nums []int) int {
	var left = make(map[int]int)
	var right = make([]int, len(nums))
	// 更新左端点对应的最远右端点
	var update = func(t, i int) {
		if l, exist := left[t]; exist {
			right[l] = i
		} else {
			left[t] = i
		}
	}
	for i, v := range nums {
		for _, fac := range f(v) {
			update(fac, i)
		}
	}
	var maxR = 0
	for l, v := range right {
		if l > maxR {
			return l - 1
		}
		maxR = max(maxR, v)
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 分解质因数
var f = func(n int) []int {
	var facs = make([]int, 0)
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			for n%i == 0 {
				n /= i
				facs = append(facs, i)
			}
		}
	}
	if n > 1 {
		facs = append(facs, n)
	}
	return facs
}
