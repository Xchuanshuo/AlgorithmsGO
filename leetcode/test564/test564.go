package test564

import (
	"math"
	"strconv"
	"strings"
)

func nearestPalindromic(str string) string {
	var n = len(str)
	var mp = map[int64]bool{int64(math.Pow10(n-1)) - 1: true,
		int64(math.Pow10(n)) + 1: true}
	target, _ := strconv.ParseInt(str, 10, 64)
	t, _ := strconv.ParseInt(str[0:(n+1)/2], 10, 64)
	for x := t - 1; x <= t+1; x++ {
		cur, _ := getNum(x, n%2 == 1)
		if cur == target {
			continue
		}
		mp[cur] = true
	}
	var res int64 = -1
	for num, _ := range mp {
		if abs(num-target) < abs(res-target) {
			res = num
		} else if abs(num-target) == abs(res-target) && num < res {
			res = num
		}
	}
	return strconv.FormatInt(res, 10)
}

func getNum(num int64, isOdd bool) (int64, error) {
	var sb strings.Builder
	sb.WriteString(strconv.FormatInt(num, 10))
	var cur = sb.String()
	var idx = len(cur) - 1
	if isOdd {
		idx = len(cur) - 2
	}
	for idx >= 0 {
		sb.WriteRune(rune(cur[idx]))
		idx--
	}
	return strconv.ParseInt(sb.String(), 10, 64)
}

func abs(num int64) int64 {
	if num < 0 {
		return -num
	}
	return num
}
