package lcw318

func maximumSubarraySum(nums []int, k int) int64 {
	var n = len(nums)
	var set = make(map[int]int)
	var total int64 = 0
	var res int64 = 0
	for i := 0; i < k-1; i++ {
		set[nums[i]]++
		total += int64(nums[i])
	}
	for i := 0; i <= n-k; i++ {
		if i > 0 {
			set[nums[i-1]]--
			if set[nums[i-1]] == 0 {
				delete(set, nums[i-1])
			}
			total -= int64(nums[i-1])
		}
		var j = i + k - 1
		set[nums[j]]++
		total += int64(nums[j])
		if len(set) == k && total > res {
			res = total
		}
	}
	return res
}
