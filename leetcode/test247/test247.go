package test247

import "fmt"

var pairs = [][]string{{"1", "1"}, {"6", "9"}, {"9", "6"}, {"8", "8"}}

func findStrobogrammatic(n int) []string {
	var cal = func(n int) []string {
		var dfs func(x int) []string
		dfs = func(x int) []string {
			if x == 0 {
				return []string{""}
			}
			if x == 1 {
				return []string{"0", "1", "8"}
			}
			var res = make([]string, 0)
			for _, s := range dfs(x - 2) {
				for _, p := range pairs {
					res = append(res, fmt.Sprintf("%s%s%s", p[0], s, p[1]))
				}
				if x != n {
					res = append(res, "0"+s+"0")
				}
			}
			return res
		}
		return dfs(n)
	}
	return cal(n)
}