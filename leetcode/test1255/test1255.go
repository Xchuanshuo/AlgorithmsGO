package test1255

/**
 * @Description https://leetcode.cn/problems/maximum-score-words-formed-by-letters/
 **/

func maxScoreWords(words []string, letters []byte, score []int) int {
	var a = make([]int, 26)
	for _, v := range letters {
		a[int(v-'a')]++
	}
	var n = len(words)
	var r = 1 << n
	var valid = make([]bool, r)
	valid[0] = true
	var res = 0
	for i := 1; i < r; i++ {
		isValid, s := true, 0
		var cur = make([]int, 26)
		for j := 0; j < n; j++ {
			if ((1 << j) & i) == 0 {
				continue
			}
			if !valid[i^(1<<j)] {
				isValid = false
				break
			}
			for _, v := range words[j] {
				var c = int(v - 'a')
				cur[c]++
				if cur[c] > a[c] {
					isValid = false
					break
				}
			}
			if !isValid {
				break
			}
		}
		if !isValid {
			continue
		}
		for k := 0; k < 26; k++ {
			s += score[k] * cur[k]
		}
		valid[i] = true
		res = max(res, s)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
