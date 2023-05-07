package test1096

import (
	"sort"
	"strings"
)

func braceExpansionII(express string) []string {
	var set = make(map[string]struct{})
	var dfs func(exp string)
	dfs = func(exp string) {
		var j = strings.Index(exp, "}")
		if j == -1 {
			set[exp] = struct{}{}
			return
		}
		var i = strings.LastIndex(exp[:j], "{")
		var ms = strings.Split(exp[i+1:j], ",")
		left, right := exp[:i], exp[j+1:]
		for _, m := range ms {
			dfs(left + m + right)
		}
	}
	dfs(express)
	var res = make([]string, 0, len(set))
	for k := range set {
		res = append(res, k)
	}
	sort.Strings(res)
	return res
}
