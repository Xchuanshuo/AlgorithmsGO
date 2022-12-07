package lcw316

/**
 * @Description https://leetcode.cn/problems/minimum-number-of-operations-to-make-arrays-similar
 * idea: 贪心 根据题意 选择递增和递减的的数对相等
		1.将nums与target进行奇偶元素分组，并从小到大排序
		2.统计递增或递减数量
 **/

import (
	"sort"
)

func makeSimilar(nums []int, target []int) int64 {
	nums1, nums2 := make([]int, 0), make([]int, 0)
	t1, t2 := make([]int, 0), make([]int, 0)
	for i, num := range nums {
		if num%2 == 1 {
			nums1 = append(nums1, num)
		} else {
			nums2 = append(nums2, num)
		}
		if target[i]%2 == 1 {
			t1 = append(t1, target[i])
		} else {
			t2 = append(t2, target[i])
		}
	}
	sort.Ints(nums1)
	sort.Ints(nums2)
	sort.Ints(t1)
	sort.Ints(t2)
	var res int64 = 0
	for i, v1 := range nums1 {
		if v1 > t1[i] {
			res += int64((v1 - t1[i]) / 2)
		}
	}
	for i, v2 := range nums2 {
		if v2 > t2[i] {
			res += int64((v2 - t2[i]) / 2)
		}
	}
	return res
}
