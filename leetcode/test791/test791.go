package test791

import "sort"

func customSortString(order string, s string) string {
	var orderMap = make(map[uint8]int)
	for i, _ := range order {
		orderMap[order[i]] = i + 1
	}
	var bytes = []byte(s)
	sort.Slice(bytes, func(i, j int) bool {
		return orderMap[bytes[i]] < orderMap[bytes[j]]
	})
	return string(bytes)
}
