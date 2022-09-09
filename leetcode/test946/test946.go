package test946

func validateStackSequences(pushed []int, popped []int) bool {
	var s = make([]int, 0)
	for _, x := range pushed {
		s = append(s, x)
		for len(s) > 0 && s[len(s)-1] == popped[0] {
			s = s[0 : len(s)-1]
			popped = popped[1:]
		}
	}
	return len(s) == 0
}
