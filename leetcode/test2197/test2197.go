package test2197

func replaceNonCoprimes(nums []int) []int {
	var s = make([]int, 0)
	for _, v := range nums {
		s = append(s, v)
		for len(s) >= 2 {
			a, b := s[len(s)-1], s[len(s)-2]
			if gcd(a, b) == 1 {
				break
			}
			s[len(s)-2] = lcm(a, b)
			s = s[0 : len(s)-1]
		}
	}
	return s
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
