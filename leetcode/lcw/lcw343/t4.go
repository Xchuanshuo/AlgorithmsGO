package lcw343

/**
 * @Description https://leetcode.cn/problems/lexicographically-smallest-beautiful-string/
 * idea: 1.逆序找到第一个满足条件的位置 2.该位置后续的元素用abc替换即可
 **/

func smallestBeautifulString(s string, k int) string {
	var bs = []byte(s)
	bs = append([]byte{'0', '0'}, bs...)
	bs = append(bs)
	var t = []byte{'a', 'b', 'c'}
	for i := len(bs) - 1; i >= 2; i-- {
		var cur = int(bs[i] - 'a' + 1)
		for ; cur < k; cur++ {
			var c = byte('a' + cur)
			if c == bs[i-1] || c == bs[i-2] {
				continue
			}
			bs[i] = c
			for j := i + 1; j < len(bs); j++ {
				for _, ch := range t {
					if j-1 >= 0 && bs[j-1] == ch {
						continue
					}
					if j-2 >= 0 && bs[j-2] == ch {
						continue
					}
					bs[j] = ch
					break
				}
			}
			return string(bs[2:])
		}

	}
	return ""
}
