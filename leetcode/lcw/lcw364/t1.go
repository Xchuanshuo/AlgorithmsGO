package lcw364

import "sort"

func maximumOddBinaryNumber(s string) string {
	var bs = []byte(s)
	sort.Slice(bs, func(i, j int) bool {
		return bs[i] > bs[j]
	})
	var t = 0
	for i := 0; i < len(bs); i++ {
		if bs[i] == '0' {
			t = i - 1
		}
	}
	bs[len(bs)-1], bs[t] = bs[t], bs[len(bs)-1]
	return string(bs)
}
