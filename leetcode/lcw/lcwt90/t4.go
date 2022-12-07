package main

/**
 * @Description https://leetcode.cn/problems/next-greater-element-iv
 * idea: 从左往右遍历 使用两个单调递减栈
 		1.s1 保证当前元素 <= 栈底元素, 若栈底元素更小 则弹出栈底元素到s2
		  此时s2存储的元素都是【右边存在一个元素比它大的】
        2.若存在当前元素大于s2中的元素，说明当前元素为【s2中元素右边第二个大于它的元素】
		  直接记录答案即可
 **/

func secondGreaterElement(nums []int) []int {
	var n = len(nums)
	var s1 = make([]int, 0)
	var s2 = make([]int, 0)
	var res = make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = -1
	}
	for i := 0; i < n; i++ {
		for len(s2) > 0 && nums[i] > nums[s2[len(s2)-1]] {
			res[s2[len(s2)-1]] = nums[i]
			s2 = s2[0 : len(s2)-1]
		}
		var t = make([]int, 0)
		for len(s1) > 0 && nums[i] > nums[s1[len(s1)-1]] {
			t = append(t, s1[len(s1)-1])
			s1 = s1[0 : len(s1)-1]
		}
		s1 = append(s1, i)
		for k := len(t) - 1; k >= 0; k-- {
			s2 = append(s2, t[k])
		}
	}
	return res
}