package test1995

func countQuadruplets(nums []int) int {
	cnts := make(map[int]int)
	res, n := 0, len(nums)
	for c := 2;c < n;c++ {
		for a := 0;a < c-1;a++ {
			cnts[nums[a] + nums[c-1]]++
		}
		for d := c+1;d < n;d++ {
			res += cnts[nums[d] - nums[c]]
		}
	}
	return res
}