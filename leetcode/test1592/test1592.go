package test1592

import "strings"

func reorderSpaces(text string) string {
	var strs = strings.Split(text, " ")
	var words = make([]string, 0)
	for _, str := range strs {
		if len(str) == 0 {
			continue
		}
		words = append(words, str)
	}
	var total = 0
	for _, v := range text {
		if v == ' ' {
			total++
		}
	}
	var cnt int
	if len(words) == 1 {
		cnt = total
	} else {
		cnt = total / (len(words) - 1)
	}
	var sb strings.Builder
	for idx, w := range words {
		sb.WriteString(w)
		if idx == len(words)-1 {
			for i := 0; i < total-cnt*(len(words)-1); i++ {
				sb.WriteString(" ")
			}
		} else {
			for i := 0; i < cnt; i++ {
				sb.WriteString(" ")
			}
		}
	}
	return sb.String()
}
