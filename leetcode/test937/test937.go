package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func reorderLogFiles(logs []string) []string {
	var spaceMap = map[string]int{}
	for _, w := range logs {
		var idx = strings.Index(w, " ")
		spaceMap[w] = idx + 1
	}
	sort.SliceStable(logs, func(i, j int) bool {
		log1, log2 := logs[i], logs[j]
		idx1, idx2 := spaceMap[log1], spaceMap[log2]
		c1, c2 := rune(log1[idx1]), rune(log2[idx2])
		if unicode.IsDigit(c1) && unicode.IsDigit(c2) {
			return false
		} else if unicode.IsDigit(c1) && !unicode.IsDigit(c2) {
			return false
		} else if !unicode.IsDigit(c1) && unicode.IsDigit(c2) {
			return true
		}
		s1, s2 := log1[idx1:], log2[idx2:]
		var val = strings.Compare(s1, s2)
		if val == 0 {
			return strings.Compare(log1[0:idx1], log2[0:idx2]) < 0
		}
		return val < 0
	})
	return logs
}

func main() {
	var strs = []string{"8 fj dzz k", "5r 446 6 3", "63 gu psub", "5 ba iedrz", "6s 87979 5", "3r 587 01", "jc 3480612", "bb wsrd kp", "b aq cojj", "r5 6316 71"}
	var res = reorderLogFiles(strs)
	fmt.Println(strings.Join(res, ","))
}
