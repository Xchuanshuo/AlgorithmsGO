package test423

import (
	"strconv"
	"strings"
)

// 统计每个字母在哪些数字中出现
func originalDigits(s string) string {
	cnts := make(map[int]int)
	for _, ch := range s {
		cnts[int(ch)]++
	}
	var cnt [10]int
	cnt[0] = cnts['z']
	cnt[2] = cnts['w']
	cnt[4] = cnts['u']
	cnt[6] = cnts['x']
	cnt[8] = cnts['g']

	cnt[3] = cnts['h'] - cnt[8]
	cnt[5] = cnts['f'] - cnt[4]
	cnt[7] = cnts['s'] - cnt[6]
	cnt[9] = cnts['i'] - cnt[5] - cnt[6] - cnt[8]
	cnt[1] = cnts['o'] - cnt[0] - cnt[2] - cnt[4]
	var sb strings.Builder
	for i := 0;i <= 9;i++ {
		for j := 0;j < cnt[i];j++ {
			sb.WriteString(strconv.Itoa(i))
		}
	}
	return sb.String()
}