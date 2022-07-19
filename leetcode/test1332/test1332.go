package test1332

func removePalindromeSub(s string) int {
	var n = len(s)
	for i := 0;i < n/2;i++ {
		if s[i] != s[n-i-1] {
			return 2
		}
	}
	return 1
}