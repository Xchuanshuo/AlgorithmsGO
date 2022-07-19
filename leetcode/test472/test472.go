package main

import (
	"fmt"
	"sort"
)

func findAllConcatenatedWordsInADict(words []string) []string {
	set := make(map[string]bool)
	for _, w := range words { set[w] = true }
	var res []string
	for _, w := range words {
		cnt, n := 0, len(w)
		dp := make([]bool, n+1)
		dp[0] = true
		for i := 1;i <= n;i++ {
			for j := i-1; j >= 0; j--{
				cur := w[j:i]
				if dp[j] && set[cur] {
					if i == n && j != 0 { cnt++ }
					dp[i] = true
					break
				}
			}
		}
		if dp[n] && cnt > 0 { res = append(res, w)}
	}
	return res
}

//----------------------------------------------------------------------------------
//----------------------------------------------------------------------------------
//----------------------------------------------------------------------------------
//----------------------------------------------------------------------------------

type Trie struct {
	children [26]*Trie
	isWorld bool
}

func (root *Trie) insert(word string) {
	node := root
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.isWorld = true
}

func (root *Trie) dfs(word string) bool {
	if word == "" {
		return true
	}
	node := root
	for i, ch := range word {
		ch -= 'a'
		node = node.children[ch]
		if node == nil { return false }
		if node.isWorld && root.dfs(word[i+1:]) {
			return true
		}
	}
	return false
}

func findAllConcatenatedWordsInADict1(words []string) []string {
	sort.Slice(words, func(i, j int) bool { return len(words[i]) < len(words[j]) })
	if words[0] == "" {
		words = words[1:]
	}
	root := &Trie{}
	res := make([]string, 0)
	for _, word := range words {
		if root.dfs(word) {
			res = append(res, word)
		} else {
			root.insert(word)
		}
	}
	return res
}

func main() {
	str := "123"
	fmt.Println(str[3:])
}
