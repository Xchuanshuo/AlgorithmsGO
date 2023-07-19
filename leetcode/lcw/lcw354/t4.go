package lcw354

func longestValidSubstring(word string, forbidden []string) int {
	var set = make(map[string]bool)
	for _, f := range forbidden {
		set[f] = true
	}
	var n = len(word)
	var res = 0
	l, r := 0, 0
	for r < n {
		for j := r; j >= 0 && j >= r-10; j-- {
			var flag = false
			for l <= r && set[word[max(l, j):r+1]] {
				l++
				flag = true
			}
			if flag {
				break
			}
		}
		res = max(res, r-l+1)
		r++
	}
	return res
}
