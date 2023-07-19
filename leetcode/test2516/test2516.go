package test2516

// 滑动窗口，右移-1, 左移+1
// 总数减去窗口内字符串不满足题意则扩大窗口
func takeCharacters(s string, k int) int {
	var n = len(s)
	var w = make([]int, 3)
	for _, v := range s {
		w[int(v-'a')]++
	}
	if w[0] < k || w[1] < k || w[2] < k {
		return -1
	}
	var res = n
	l, r := -1, 0
	for ; r < n; r++ {
		var c = int(s[r] - 'a')
		w[c]--
		for l < r && w[c] < k {
			l++
			w[int(s[l]-'a')]++
		}
		res = min(res, n-(r-l))
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
