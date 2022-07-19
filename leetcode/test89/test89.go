package test89

func grayCode(n int) []int {
	data := make([]int, 0)
	for i := 0;i < 1<<n;i++ {
		data = append(data, i^(i>>1))
	}
	return data
}