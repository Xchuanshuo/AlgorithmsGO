package test6

import "strings"

func convert(s string, numRows int) string {
	var rows = make([][]rune, 0)
	for i := 0; i < numRows; i++ {
		rows = append(rows, make([]rune, 0))
	}
	idx, flag := 0, -1
	for _, v := range s {
		rows[idx] = append(rows[idx], v)
		if idx == 0 || idx == numRows-1 {
			flag = -flag
		}
		idx += flag
	}
	var sb strings.Builder
	for _, r := range rows {
		sb.WriteString(string(r))
	}
	return sb.String()
}
