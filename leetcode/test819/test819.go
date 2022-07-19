package test819

import (
	"sort"
	"strings"
)

func mostCommonWord(paragraph string, banned []string) string {
	var t = make(map[string]bool)
	for _, s := range banned {
		t[s] = true
	}
	paragraph = strings.ReplaceAll(paragraph, ",", " ")
	paragraph = strings.ReplaceAll(paragraph, ".", " ")
	paragraph = strings.ReplaceAll(paragraph, "?", " ")
	paragraph = strings.ReplaceAll(paragraph, ";", " ")
	paragraph = strings.ReplaceAll(paragraph, "'", " ")
	paragraph = strings.ReplaceAll(paragraph, "!", " ")
	var strs = strings.Split(paragraph, " ")
	var countMap = make(map[string]int)
	var arr = make([]string, 0)
	for _, s := range strs {
		if len(s) == 0 {
			continue
		}
		s = strings.ToLower(s)
		if _, exist := countMap[s]; !exist {
			arr = append(arr, s)
		}
		countMap[s]++
	}
	sort.Slice(arr, func(i, j int) bool {
		return countMap[arr[i]] > countMap[arr[j]]
	})
	for _, a := range arr {
		if _, exist := t[a]; !exist {
			return a
		}
	}
	return ""
}
