package main

import (
	"strconv"
	"strings"
)

func exclusiveTime(n int, logs []string) []int {
	var res = make([]int, n)
	var s = make([]*Log, 0)
	for i := 0; i < len(logs); i++ {
		var cur = strings.Split(logs[i], ":")
		id, _ := strconv.Atoi(cur[0])
		time, _ := strconv.Atoi(cur[2])
		if cur[1] == "start" {
			s = append(s, &Log{fid: id, time: time})
		} else {
			var peek = s[len(s)-1]
			s = s[0 : len(s)-1]
			var offset = time - peek.time + 1
			res[id] += offset
			if len(s) > 0 {
				res[s[len(s)-1].fid] -= offset
			}
		}
	}
	return res
}

type Log struct {
	fid  int
	time int
}
