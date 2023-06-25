package test843

import (
	"sort"
)

/**
 * @Description https://leetcode.cn/problems/guess-the-word/
 * idea: 题意: 给定一个单词组，限定猜测次数的情况下能否猜测到正确单车
		猜测策略: 由于每次猜测返回res，不正确进行下一轮猜测时，能根据当前的结果排除掉数量!=res的字符串
				所以需要猜测时尽量选择【可以排除更多单词】的目标词汇，用cnts记录与任意单词不同数量的出现次数
				若最大数量出现的次数小，则说明不同数量分布更平均，优先选这个即可
 **/

func findSecretWord(words []string, master *Master) {
	var dfs func(words []string)
	dfs = func(words []string) {
		var cur = selectOne(words)
		var res = master.Guess(cur)
		if res == 6 {
			return
		}
		var nws = make([]string, 0)
		for _, w := range words {
			if getCnt(cur, w) != res {
				continue
			}
			nws = append(nws, w)
		}
		dfs(nws)
	}
	dfs(words)
}

func selectOne(words []string) string {
	t, min := 0, len(words)
	for i, w := range words {
		var cnts = make([]int, 6)
		for j, v := range words {
			if j == i {
				continue
			}
			cnts[getCnt(w, v)]++
		}
		sort.Ints(cnts)
		if cnts[5] < min {
			min = cnts[5]
			t = i
		}
	}
	return words[t]
}

func getCnt(s, t string) int {
	var cnt = 0
	for i := range s {
		if s[i] == t[i] {
			cnt++
		}
	}
	return cnt
}

type Master struct {
}

func (this *Master) Guess(word string) int {
	return 0
}
