package test735

func asteroidCollision(asteroids []int) []int {
	var s = make([]int, 0)
	for _, a := range asteroids {
		var flag = true
		for len(s) > 0 && s[len(s)-1] > 0 && a < 0 {
			if s[len(s)-1] > -a {
				flag = false
				break
			} else if s[len(s)-1] < -a {
				s = s[0 : len(s)-1]
				flag = true
			} else {
				s = s[0 : len(s)-1]
				flag = false
				break
			}
		}
		if flag {
			s = append(s, a)
		}
	}
	return s
}
