package test71

import (
	"strings"
)

func simplifyPath(path string) string {
	strs := strings.Split(path, "/")
	q := make([]string, 0)
	for _, str := range strs {
		if str == ".." {
			if len(q) > 0 {q = q[0:len(q)-1]}
		} else if len(str)>0 && str != "." {
			q = append(q, str)
		}
	}
	var sb strings.Builder
	if len(q) == 0 {
		sb.WriteString("/")
	} else {
		for len(q) > 0 {
			sb.WriteString("/")
			sb.WriteString(q[0])
			q = q[1:]
		}
	}
	return sb.String()
}