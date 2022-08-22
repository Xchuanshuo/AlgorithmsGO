package main

import "fmt"

/**
 * @Description https://leetcode.cn/problems/check-if-string-is-transformable-with-substring-sort-operations/
 * idea: 1.模拟冒泡排序，对连续多个字符排序，都可以转换成连续两个字符交换
 *       2.能否交换需要满足：当前字符 【左边】的字符都【大于】该字符，右边的都【小于】该字符
 *		 3.每个目标字符t[i]在源字符串s中选取最早出现的那个
 **/

func isTransformable(s string, t string) bool {
	var posMap = make(map[int][]int)
	for idx, v := range s {
		var p = int(v - '0')
		posMap[p] = append(posMap[p], idx)
	}
	for _, v := range t {
		var p = int(v - '0')
		if len(posMap[p]) == 0 {
			return false
		}
		for i := 0; i < p; i++ {
			if len(posMap[i]) > 0 && posMap[i][0] < posMap[p][0] {
				return false
			}
		}
		posMap[p] = posMap[p][1:]
	}
	return true
}
func main() {
	var s = "12345"
	var t = "12435"
	res := isTransformable(s, t)
	fmt.Println(res)
}
