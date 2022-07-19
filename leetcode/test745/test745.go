package test745

type WordFilter struct {
	PreRoot *Trie
	SufRoot *Trie
}

func Constructor(words []string) WordFilter {
	var preRoot = NewTrie()
	var sufRoot = NewTrie()
	for idx, w := range words {
		preRoot.insert(w, idx, false)
		sufRoot.insert(w, idx, true)
	}
	return WordFilter{
		PreRoot: preRoot, SufRoot: sufRoot,
	}
}

func (this *WordFilter) F(pref string, suff string) int {
	var a1 = this.PreRoot.search(pref, false)
	var a2 = this.SufRoot.search(suff, true)
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}
	i, j := len(a1)-1, len(a2)-1
	for i >= 0 && j >= 0 {
		c1, c2 := a1[i], a2[j]
		if c1 == c2 {
			return c1
		} else if c1 > c2 {
			i--
		} else {
			j--
		}
	}
	return -1
}

type Trie struct {
	Root *Node
}

func NewTrie() *Trie {
	return &Trie{Root: NewNode()}
}

func (this *Trie) insert(word string, idx int, isReverse bool) {
	var root = this.Root
	var n = len(word)
	for i := 0; i < len(word); i++ {
		var w = word[i]
		if isReverse {
			w = word[n-i-1]
		}
		if _, exist := root.children[w]; !exist {
			root.children[w] = NewNode()
		}
		root = root.children[w]
		root.idxs = append(root.idxs, idx)
	}
}

func (this *Trie) search(word string, isReverse bool) []int {
	var root = this.Root
	var n = len(word)
	for i := 0; i < len(word); i++ {
		var w = word[i]
		if isReverse {
			w = word[n-i-1]
		}
		if _, exist := root.children[w]; !exist {
			return []int{}
		}
		root = root.children[w]
	}
	return root.idxs
}

type Node struct {
	idxs     []int
	children map[byte]*Node
}

func NewNode() *Node {
	return &Node{idxs: []int{}, children: map[byte]*Node{}}
}
