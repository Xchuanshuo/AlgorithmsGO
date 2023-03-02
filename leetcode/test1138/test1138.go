package test1138

import (
	"strings"
)

var g = []string{"abcde", "fghij", "klmno", "pqrst", "uvwxy", "z####"}
var idxMap = make(map[uint8][]int)

func init() {
	for i := 0; i < 6; i++ {
		for j := 0; j < 5; j++ {
			idxMap[g[i][j]] = []int{i, j}
		}
	}
}

func alphabetBoardPath(target string) string {
	var bfs = func(s, t uint8) string {
		sx, sy := idxMap[s][0], idxMap[s][1]
		tx, ty := idxMap[t][0], idxMap[t][1]
		if sx == tx && sy == ty {
			return "!"
		}
		var b = make([]byte, 0)
		for sx != tx || sy != ty {
			if sx < tx && g[sx+1][sy] != '#' {
				b = append(b, 'D')
				sx++
			} else if sx > tx && g[sx-1][sy] != '#' {
				b = append(b, 'U')
				sx--
			} else if sy < ty && g[sx][sy+1] != '#' {
				b = append(b, 'R')
				sy++
			} else if sy > ty && g[sx][sy-1] != '#' {
				b = append(b, 'L')
				sy--
			}
		}
		b = append(b, '!')
		return string(b)
	}
	var sb strings.Builder
	var pre uint8 = 'a'
	for i := 0; i < len(target); i++ {
		sb.WriteString(bfs(pre, target[i]))
		pre = target[i]
	}
	return sb.String()
}
