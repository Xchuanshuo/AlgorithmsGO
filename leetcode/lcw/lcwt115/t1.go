package lcwt115

import "strconv"

func lastVisitedIntegers(words []string) []int {
	var nums = make([]int, 0)
	var res = make([]int, 0)
	var k = 0
	for _, s := range words {
		if s == "prev" {
			k++
			if k > len(nums) {
				res = append(res, -1)
			} else {
				res = append(res, nums[len(nums)-1-(k-1)])
			}
		} else {
			v, _ := strconv.Atoi(s)
			nums = append(nums, v)
			k = 0
		}
	}
	return res
}
