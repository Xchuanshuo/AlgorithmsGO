package lcw337

func evenOddBit(n int) []int {
	var a = make([]int, 0)
	for n != 0 {
		a = append(a, n%2)
		n /= 2
	}
	odd, even := 0, 0
	for i := len(a)-1;i >= 0;i-- {
		if (n-i-1)%2 == 0 {
			if a[i] == 1 { odd++ }
		} else {
			if a[i] == 1 { even++ }
		}
	}
	return []int{even, odd}
}
