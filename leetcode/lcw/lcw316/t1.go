package lcw316

import (
	"strconv"
	"strings"
)

func haveConflict(event1 []string, event2 []string) bool {
	l1, r1 := getV(event1[0]), getV(event1[1])
	l2, r2 := getV(event2[0]), getV(event2[1])
	return !(l2 > r1 || l1 > r2)
}

func getV(s string) int {
	var strs = strings.Split(s, ":")
	v1, _ := strconv.Atoi(strs[0])
	v2, _ := strconv.Atoi(strs[1])
	return v1*60 + v2
}
