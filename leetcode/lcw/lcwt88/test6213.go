package main

func xorAllNums(nums1 []int, nums2 []int) int {
	n1, n2 := len(nums1), len(nums2)
	if n1%2 == 0 && n2%2 == 0 {
		return 0
	}
	var res = 0
	if n1%2 == 0 {
		for _, num := range nums1 {
			res ^= num
		}
	} else if n2%2 == 0 {
		for _, num := range nums2 {
			res ^= num
		}
	} else {
		for _, num := range nums1 {
			res ^= num
		}
		for _, num := range nums2 {
			res ^= num
		}
	}
	return res
}
