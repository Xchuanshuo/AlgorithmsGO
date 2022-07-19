package test648

import "strings"

func replaceWords(dictionary []string, sentence string) string {
	type tire map[rune]tire
	root := &Node{Children: map[rune]*Node{}}
	for _, d := range dictionary {
		cur := root
		for _, c := range d {
			if cur.Children[c] == nil {
				cur.Children[c] = &Node{Children: map[rune]*Node{}}
			}
			cur = cur.Children[c]
		}
		cur.IsWord = true
	}
	var words = strings.Split(sentence, " ")
	for i, word := range words {
		cur := root
		for j, c := range word {
			if cur == nil {
				break
			}
			if cur.IsWord {
				words[i] = word[:j]
				break
			}
			cur = cur.Children[c]
		}
	}
	return strings.Join(words, " ")
}

type Node struct {
	Children map[rune]*Node
	IsWord   bool
}
