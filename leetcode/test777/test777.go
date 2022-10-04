package test777

import "strings"

func canTransform(start string, end string) bool {
	if strings.ReplaceAll(start, "X", "") != strings.ReplaceAll(end, "X", "") {
		return false
	}
	var n = len(start)
	i, j := 0, 0
	for i < n && j < n {
		if start[i] == 'X' {
			i++
			continue
		}
		var cj = end[j]
		for j < n && end[j] == 'X' {
			j++
			if j < n {
				cj = end[j]
			}
		}
		if start[i] != cj {
			return false
		}
		if start[i] == 'L' && i < j {
			return false
		}
		if start[i] == 'R' && i > j {
			return false
		}
		i, j = i+1, j+1
	}
	for i < n && start[i] == 'X' {
		i++
	}
	for j < n && end[j] == 'X' {
		j++
	}
	return i == j
}
