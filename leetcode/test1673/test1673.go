package test1673

func mostCompetitive(nums []int, k int) []int {
	var n = len(nums)
	var s = make([]int, 0)
	for i, v := range nums {
		for len(s) > 0 && nums[i] < s[len(s)-1] && len(s)-1+(n-i) >= k {
			s = s[0 : len(s)-1]
		}
		s = append(s, v)
	}
	return s[0:k]
}
