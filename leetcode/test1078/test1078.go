package test1078

import "strings"

func findOcurrences(text string, first string, second string) []string {
	strs := strings.Split(text, " ")
	res := make([]string, 0)
	for i := 0;i < len(strs)-2;i++ {
		if strs[i] == first && strs[i+1] == second {
			res = append(res, strs[i+2])
		}
	}
	return res
}