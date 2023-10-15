package lcw363

/**
 * @Description https://leetcode.cn/problems/maximum-element-sum-of-a-complete-subset-of-indices/
 * idea: 一个数对乘积为完全平方数 => 每个数去除完全平方因子后相等, 所以按去完全平方因子后分组分别计算最大和即可
		 时间复杂度为 n*sqrt(n)
 **/

import "math"

func maximumSum(nums []int) int64 {
	var mp = make(map[int][]int)
	for i, _ := range nums {
		var x = cal(i + 1)
		mp[x] = append(mp[x], i)
	}
	var res = int64(0)
	for _, vs := range mp {
		var cur int64
		for _, i := range vs {
			cur += int64(nums[i])
		}
		res = maxI64(res, cur)
	}
	return res
}

// 对于指定数去除完全平方数后的因子
var cal = func(t int) int {
	for i := int(math.Sqrt(float64(t))); i >= 2; i-- {
		if t%(i*i) == 0 {
			return t / (i * i)
		}
	}
	return t
}

func maxI64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
