package test1342

func numberOfSteps(num int) int {
	res := 0
	for num != 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num -= 1
		}
		res++
	}
	return res
}
