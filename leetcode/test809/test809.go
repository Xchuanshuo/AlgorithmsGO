package test809

func expressiveWords(s string, words []string) int {
	var srcArr = make([][]int, 0)
	var cnt = 0
	for i, c := range s {
		cnt++
		if i == len(s)-1 || s[i] != s[i+1] {
			srcArr = append(srcArr, []int{int(c), cnt})
			cnt = 0
		}
	}
	var res = 0
	for _, word := range words {
		var tArr = make([][]int, 0)
		cnt = 0
		for i, c := range word {
			cnt++
			if i == len(word)-1 || word[i] != word[i+1] {
				tArr = append(tArr, []int{int(c), cnt})
				cnt = 0
			}
		}
		if len(tArr) != len(srcArr) {
			continue
		}
		var isValid = true
		for i, c := range tArr {

			if c[0] != srcArr[i][0] || srcArr[i][1] < tArr[i][1] ||
				(srcArr[i][1] > tArr[i][1] && srcArr[i][1] < 3) {
				isValid = false
				break
			}
		}
		if isValid {
			res++
		}
	}
	return res
}
