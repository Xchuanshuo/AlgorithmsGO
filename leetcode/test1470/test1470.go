package test1470

func shuffle(nums []int, n int) []int {
	i, j := 0, n
	var res = make([]int, 0)
	var flag = true
	for i < n || j < 2*n {
		if flag {
			res = append(res, nums[i])
			i++
		} else {
			res = append(res, nums[j])
			j++
		}
		flag = !flag
	}
	return res
}
