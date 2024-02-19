package test421

func findMaximumXOR(nums []int) int {
	var trie = NewTrie()
	for _, v := range nums {
		trie.insert(v)
	}
	var res = 0
	for _, v := range nums {
		res = max(res, trie.getXor(v))
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
	}
	cur.isWord = true
}

func (t *Trie) getXor(num int) int {
	var cur = t.root
	var val = 0
	for i := 31; i >= 0; i-- {
		var tb = (num >> i & 1) ^ 1
		if cur.children[tb] != nil {
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
	children [2]*Node
}
