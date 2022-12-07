package main

import (
	"fmt"
	"strings"
)

func oddString(words []string) string {
	var mp = make(map[string]int)
	for _, w := range words {
		var s strings.Builder
		for i := 1; i < len(w); i++ {
			s.WriteString(fmt.Sprintf("%d_", w[i]-w[i-1]))
		}
		mp[s.String()]++
	}
	for _, w := range words {
		var s strings.Builder
		for i := 1; i < len(w); i++ {
			s.WriteString(fmt.Sprintf("%d_", w[i]-w[i-1]))
		}
		if mp[s.String()] == 1 {
			return w
		}
	}
	return ""
}
