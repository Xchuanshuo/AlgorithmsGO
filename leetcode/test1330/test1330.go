package test1330

import "math"

/**
 * @Description https://leetcode.cn/problems/reverse-subarray-to-maximize-array-value/
 * idea: 1.观察规律，翻转后中间元素的绝对值不变，两边值变
         2.列出翻转前后的边界绝对值式子, 假设左边界为a[i+1], 右边界为b[i]
		   那么只需要计算翻转前后的变化量即可，翻转后的值 = 变化量 + 不进行翻转的原始数组值，即
		  |a[i]-a[i+1]|+|b[i]-b[i+1]| - |a[i]-b[i]|-|a[i+1]-b[i+1]|
		 3.涉及到a[i],a[i+1],b[i],b[i+1]4个变量，且数据范围为3e4, 考虑将绝对值式子进行化简，
		  分情况讨论a,b的大小，最终可以得到关于各自计算a,b两部分，形如【nums[j] + nums[i]】 的式子
		  这样每访问到一个元素，记录到当前为止得到的最大nums[j]即可
		 4.遍历数组过程中计算翻转当前位置得到的最大变化量即可
 **/

var INF = math.MaxInt32 / 2

func maxValueAfterReverse(nums []int) int {
	var n = len(nums)
	a, b := nums[0], nums[n-1]
	sum, dif := 0, 0
	var pre = make([]int, 4)
	for i := 0; i < 4; i++ {
		pre[i] = INF
	}
	for i := 0; i < n-1; i++ {
		x, y := nums[i], nums[i+1]
		var d = abs(x - y)
		sum += d
		dif = minA([]int{
			dif, d - abs(a-y), d - abs(x-b),
			d + x + y + pre[0],
			d + x - y + pre[1],
			d - x + y + pre[2],
			d - x - y + pre[3],
		})
		pre[0] = min(pre[0], d-x-y)
		pre[1] = min(pre[1], d-x+y)
		pre[2] = min(pre[2], d+x-y)
		pre[3] = min(pre[3], d+x+y)
	}
	return sum - dif
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minA(a []int) int {
	var res = math.MaxInt32
	for _, v := range a {
		if v < res {
			res = v
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
