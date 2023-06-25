package test880

/**
 * @Description https://leetcode.cn/problems/decoded-string-at-index/
 * idea: 直接模拟会超内存，考虑当前总长度l与k的关系，有以下几种情况
		 1.l<k，继续看后续的字符，累加总长度
		 2.l==k或 l>k且k%l==0 当前字符串结尾即为答案
		 3.l>k, 且不能对k整除，此时将l重复n次后，k指向的位置等同于在未重复之前k%l指向的位置
		   直接当作子问题计算即可
		 整体时间复杂度O(n^2)
 **/
import "fmt"

func decodeAtIndex(s string, t int) string {
	var k = int64(t)
	var l = int64(0)
	for i, v := range s {
		var add int64
		if v >= '0' && v <= '9' {
			add = l * int64(v-'1')
		} else {
			add = 1
		}
		if l+add < k {
			l += add
			continue
		}
		if l+add == k || k%l == 0 {
			// 当前k刚好指向l重复n次后的结尾
			for j := i; j >= 0; j-- {
				if s[j] >= 'a' && s[j] <= 'z' {
					return fmt.Sprintf("%c", s[j])
				}
			}
		} else {
			// 寻找子问题 k的原位置指向字符与l重复n次后k%l的位置相同
			return decodeAtIndex(s[0:i], int(k%l))
		}
	}
	return ""
}
