package lcw355

import "sort"

/**
 * @Description https://leetcode.cn/problems/maximum-number-of-groups-with-increasing-length/
 * idea: 1.构造数组与具体的数无关, 频率更高的数能参与更多group的构建，所以考虑按频率排序
		 2.如何去重? 对于新加入的数字，若导致剩余总数>group，首先考虑将原来group组中的每个组剔除掉一个元素，
		  然后用当前剩余的总数去替换原来位置，前面频率更小的数字可以构造group个组，当前加入更大的后也【一定】可以构造。
		  而对于剔除的k个数，因为前面每组元素不一样，且比前一组多一个元素，所以【一定】能找出k个不同的数
		  对于第group+1组，使用剔除的k个数 + 当前预留的新数进行构造，多余的数加入缓存。
		 3.考虑排序方式, 假设按照频率递减序，则按照2所诉方式构建则会导致重复，所以按递增顺序即可
 **/

func maxIncreasingGroups(nums []int) int {
	sort.Ints(nums)
	total, group := 0, 0
	for _, v := range nums {
		total += v
		if total > group {
			group++
			total -= group
		}
	}
	return group
}
