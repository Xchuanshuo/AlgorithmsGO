package test1464

func maxProduct(nums []int) int {
	f, s := 0, 0
	for _, n := range nums {
		if n > f {
			s = f
			f = n
		} else if n > s {
			s = n
		}
	}
	return (f - 1) * (s - 1)
}
