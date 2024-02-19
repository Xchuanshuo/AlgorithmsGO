package lcw369

func minSum(nums1 []int, nums2 []int) int64 {
	var s1, s2 int64
	var z1, z2 = 0, 0
	for _, v := range nums1 {
		s1 += int64(v)
		if v == 0 {
			z1++
		}
	}
	for _, v := range nums2 {
		s2 += int64(v)
		if v == 0 {
			z2++
		}
	}
	if z1 == 0 && z2 == 0 && s1 != s2 {
		return -1
	}
	if (z1 == 0 && (s1 < s2 || s2+int64(z2) > s1)) || (z2 == 0 && (s2 < s1 || s1+int64(z1) > s2)) {
		return -1
	}
	return maxI64(s2+int64(z2), s1+int64(z1))
}

func maxI64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
