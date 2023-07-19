package s2

import "strings"

type Encrypter struct {
	enMap map[byte]string
	deMap map[string][]byte
	trie  *Trie
}

func Constructor(keys []byte, values []string, dictionary []string) Encrypter {
	var enMap = make(map[byte]string)
	var deMap = make(map[string][]byte)
	for i, v := range keys {
		enMap[v] = values[i]
		deMap[values[i]] = append(deMap[values[i]], v)
	}
	var trie = NewTrie()
	for _, d := range dictionary {
		trie.Insert(d)
	}
	return Encrypter{enMap, deMap, trie}
}

func (this *Encrypter) Encrypt(word1 string) string {
	var bs = make([]byte, 0)
	for _, c := range word1 {
		bs = append(bs, this.enMap[byte(c)]...)
	}
	return string(bs)
}

func (this *Encrypter) Decrypt(word2 string) int {
	var dfs func(*Node, string) int
	dfs = func(root *Node, word string) int {
		if len(word) == 0 {
			if root.isWord {
				return 1
			}
			return 0
		}
		var res = 0
		var bs = this.deMap[word[0:2]]
		for _, b := range bs {
			if root.next[b] == nil {
				continue
			}
			res += dfs(root.next[b], word[2:])
		}
		return res
	}
	return dfs(this.trie.root, word2)
}

type Node struct {
	isWord bool
	prefix string
	next   map[byte]*Node
}

func newNode(isWord bool, prefix string) *Node {
	return &Node{isWord: isWord, prefix: prefix, next: make(map[byte]*Node)}
}

type Trie struct {
	root *Node
	size int
}

func NewTrie() *Trie {
	return &Trie{root: newNode(false, ""), size: 0}
}

func (t *Trie) GetSize() int {
	return t.size
}

func (t *Trie) Insert(word string) {
	cur := t.root
	var sb strings.Builder
	for i := 0; i < len(word); i++ {
		c := word[i]
		sb.WriteByte(c)
		if _, ok := cur.next[c]; !ok {
			cur.next[c] = newNode(false, sb.String())
		}
		cur = cur.next[c]
	}
	if !cur.isWord {
		cur.isWord = true
		t.size++
	}
}

func (t *Trie) Search(word string) bool {
	cur := t.root
	node := findPrefixNode(word, cur)
	return node != nil && node.isWord
}

func (t *Trie) SearchPrefix(prefix string) bool {
	cur := t.root
	return findPrefixNode(prefix, cur) != nil
}

func findPrefixNode(prefix string, cur *Node) *Node {
	for i := 0; i < len(prefix); i++ {
		c := prefix[i]
		if _, ok := cur.next[c]; !ok {
			return nil
		}
		cur = cur.next[c]
	}
	return cur
}
