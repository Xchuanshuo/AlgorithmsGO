package test1702

/**
 * @Description https://leetcode.cn/problems/maximum-binary-string-after-change/
 * idea: 贪心, 优先让高位变为1。 根据构造用例可以发现，任何后面的0都可以移到1前面, 如"10" => "01",
		对于连续的0, 可以变为10, 所以对于1开始的前缀无需移动, 去除1前缀后的部分所有0移动到1前面，再从第
		一个位置开始依次构造1
 **/

import "sort"

func maximumBinaryString(binary string) string {
	var bs = []byte(binary)
	var idx = -1
	if bs[0] == '1' {
		for i := 0; i < len(bs); i++ {
			if bs[i] != '1' {
				break
			}
			idx = i
		}
	}
	var l = string(bs[0 : idx+1])
	bs = bs[idx+1:]
	sort.Slice(bs, func(i, j int) bool {
		return bs[i] < bs[j]
	})
	for i := 1; i < len(bs); i++ {
		if string(bs[i-1:i+1]) == "00" {
			bs[i-1] = '1'
		}
	}
	return l + string(bs)
}