package test828

/**
 * @Description https://leetcode.cn/problems/count-unique-characters-of-all-substrings-of-a-given-string
 * idea:计算每个字符产生的贡献, 即只出现一次时能给多少子字符串产生贡献
 **/

func uniqueLetterString(s string) int {
	var pos = make([][]int, 26)
	for i, v := range s {
		var p = int(v - 'A')
		if len(pos[p]) == 0 {
			pos[p] = append(pos[p], -1)
		}
		pos[p] = append(pos[p], i)
	}
	var res = 0
	for i := 0; i < 26; i++ {
		pos[i] = append(pos[i], len(s))
		for j := 1; j < len(pos[i])-1; j++ {
			res += (pos[i][j] - pos[i][j-1]) * (pos[i][j+1] - pos[i][j])
		}
	}
	return res
}