package lcw371

/**
 * @Description https://leetcode.cn/problems/maximum-strong-pair-xor-ii/
 * idea: 可撤销字典树 + 滑动窗口.
		1.排序, x为之前的数，y为后面的数，公式化简为 y-x <= x, 即 2x >= y
		2.使用字典树求最大异或值, 维护窗口左边界，对字典树进行单词删除操作
 **/

import "sort"

func maximumStrongPairXor(nums []int) int {
	var res = 0
	sort.Ints(nums)
	var trie = NewTrie()
	var l = 0
	for r, v := range nums {
		var n = r + 1
		var t = sort.Search(n, func(i int) bool {
			return 2*nums[i] >= v
		})
		for ; l < t; l++ {
			trie.delete(nums[l])
		}
		trie.insert(v)
		res = max(res, trie.getXor(v))
	}
	return res
}

type Trie struct {
	root *Node
}

func NewTrie() Trie {
	return Trie{root: &Node{}}
}

func (t *Trie) insert(num int) {
	var cur = t.root
	for i := 31; i >= 0; i-- {
		var b = num >> i & 1
		if cur.children[b] == nil {
			cur.children[b] = &Node{}
		}
		cur = cur.children[b]
		cur.cnt++
	}
	cur.isWord = true
}

// trie删除单词
func (t *Trie) delete(num int) {
	var cur = t.root
	for i := 31; i >= 0; i-- {
		var b = num >> i & 1
		if cur.children[b] == nil {
			cur.children[b] = &Node{}
		}
		cur = cur.children[b]
		cur.cnt--
	}
	cur.isWord = false
}

func (t *Trie) getXor(num int) int {
	var cur = t.root
	var val = 0
	for i := 31; i >= 0; i-- {
		var tb = (num >> i & 1) ^ 1
		if cur.children[tb] != nil && cur.children[tb].cnt > 0 {
			cur = cur.children[tb]
			val |= 1 << i
		} else {
			cur = cur.children[tb^1]
		}
	}
	return val
}

type Node struct {
	isWord   bool
	cnt      int // 记录使用单词的前缀个数
	children [2]*Node
}
