package test1694

import "strings"

func reformatNumber(number string) string {
	number = strings.ReplaceAll(number, "-", "")
	number = strings.ReplaceAll(number, " ", "")
	var res = make([]byte, 0)
	var remain = len(number) % 3
	if remain == 1 {
		remain += 3
	}
	for i := 0; i < len(number)-remain; i++ {
		if i != 0 && i%3 == 0 {
			res = append(res, '-')
		}
		res = append(res, number[i])
	}
	if len(res) > 0 && remain != 0 {
		res = append(res, '-')
	}
	var s = len(number) - remain
	if remain == 3 || remain == 2 {
		res = append(res, number[len(number)-remain:]...)
	} else if remain == 4 {
		res = append(res, number[s:s+2]...)
		res = append(res, '-')
		res = append(res, number[s+2:]...)
	}
	return string(res)
}
