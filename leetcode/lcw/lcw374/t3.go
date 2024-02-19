package lcw374

func countCompleteSubstrings(word string, k int) int {
	var wind = make([]int, 26)
	l, r := 0, len(word)
	var res = 0
	for r < len(wind) {
		var p = int(word[r] - 'a')
		wind[p]++
		for wind[p] > k {
			p = int(word[l] - 'a')
			wind[p]--
		}
		if r > 0 && int(word[r])-int(word[r-1]) > 2 {
			l = r
		}
		var valid = true
		for i := 0; i < 26; i++ {
			if wind[i] > 0 && wind[i] != k {
				valid = false
				break
			}
		}
		if valid {
			res++
		}
		r++
	}
	return res
}
