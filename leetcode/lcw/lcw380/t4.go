package main

/**
 * @Description https://leetcode.cn/problems/find-beautiful-indices-in-the-given-array-ii
 * idea: kmp + 双指针
		1.kmp查找模式串a、b的所有出现位置idxA、idxB
		2.双指针枚举索引idxA和idxB，索引B的当前位置大于索引A后则停止移动
 **/

import (
	"fmt"
)

func beautifulIndices(s string, a string, b string, k int) []int {
	var idxA = kmps(s, a)
	var idxB = kmps(s, b)
	var res = make([]int, 0)
	var bi = 0
	for _, ai := range idxA {
		if bi >= len(idxB) {
			break
		}
		for bi < len(idxB) && idxB[bi] < ai && abs(ai-idxB[bi]) > k {
			bi++
		}
		if bi < len(idxB) && abs(ai-idxB[bi]) <= k {
			res = append(res, ai)
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// kmp查找所有匹配字符串的索引初始位置
func kmps(master string, pattern string) []int {
	// next数组表示最长公共前后缀子串中 '前缀子串' 的最后一个字符的位置
	var next = getNext(pattern)
	ml, pl := len(master), len(pattern)
	var res = make([]int, 0)
	var j = 0
	for i := 0; i < ml; i++ {
		for j > 0 && master[i] != pattern[j] {
			j = next[j-1] + 1
		}
		if master[i] == pattern[j] {
			j++
		}
		if j == pl {
			res = append(res, i-pl+1)
			j = next[j-1] + 1
		}
	}
	return res
}

// next表
func getNext(pattern string) []int {
	next := make([]int, len(pattern))
	next[0] = -1
	k := -1
	for i := 1; i < len(pattern); i++ {
		for k != -1 && pattern[k+1] != pattern[i] {
			// 当前位置字符与已经匹配的前缀的下一个字符不相等 则去前一个位置寻找
			k = next[k]
		}
		if pattern[k+1] == pattern[i] {
			k++
		}
		next[i] = k
	}
	return next
}

func main() {
	var s = "isawsquirrelnearmysquirrelhouseohmy"
	var a = "my"
	var b = "squirrel"
	var k = 15
	var res = beautifulIndices(s, a, b, k)
	fmt.Println(res)
}
