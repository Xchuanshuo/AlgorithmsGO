package main

import (
	"fmt"
)

/**
 * @Description https://leetcode.cn/contest/weekly-contest-313/problems/maximum-deletions-on-a-string/
 * idea: 1.dfs+记忆化搜索 2.字符串hash + dp, dp[i]表示从第i个字符串开始删除需要的最大删除数
 **/

func deleteString(s string) int {
	var hash = NewStrHash(s)
	var n = len(s)
	var dp = make([]int, n+1)
	for i := n; i > 0; i-- {
		dp[i] = 1
		for j := i; j < i+(n-i+1)/2; j++ {
			h1, h2 := hash.Get(i, j), hash.Get(j+1, j+1+j-i)
			if h1 == h2 {
				dp[i] = max(dp[i], dp[j+1]+1)
			}
		}
	}
	return dp[1]
}

type StrHash struct {
	h, p []int
}

func NewStrHash(s string) StrHash {
	var n = len(s)
	h, p := make([]int, n+1), make([]int, n+1)
	p[0] = 1
	for i := 0; i < n; i++ {
		h[i+1] = h[i]*P + int(s[i])
		p[i+1] = p[i] * P
	}
	return StrHash{h, p}
}

const P = 131

func (s StrHash) Get(l, r int) int {
	return s.h[r] - s.h[l-1]*s.p[r-l+1]
}

func deleteString1(s string) int {
	var dfs func(s string) int
	var mp = make(map[string]int)
	dfs = func(s string) int {
		if v, exist := mp[s]; exist {
			return v
		}
		var res = 0
		var isValid = false
		for i := 1; i <= len(s)/2; i++ {
			if s[0:i] == s[i:2*i] {
				isValid = true
				res = max(res, 1+dfs(s[i:]))
			}
		}
		if !isValid {
			res = 1
		}
		mp[s] = res
		return res
	}
	return dfs(s)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	var s = "zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz"
	for i := 5; i >= 0; i-- {
		s += s
	}
	fmt.Println(deleteString(s))
}
