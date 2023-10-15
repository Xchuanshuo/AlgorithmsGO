package test2512

import (
	"sort"
	"strings"
)

func topStudents(positive_feedback []string, negative_feedback []string, report []string, student_id []int, k int) []int {
	var scoreMap = make(map[string]int)
	for _, s := range positive_feedback {
		scoreMap[s] = 3
	}
	for _, s := range negative_feedback {
		scoreMap[s] = -1
	}
	var ids = make([]int, 0)
	var idToScore = make(map[int]int)
	for i, r := range report {
		var ss = strings.Split(r, " ")
		var score = 0
		for _, s := range ss {
			score += scoreMap[s]
		}
		idToScore[student_id[i]] = score
		ids = append(ids, student_id[i])
	}
	sort.Slice(ids, func(i, j int) bool {
		if idToScore[ids[i]] != idToScore[ids[j]] {
			return idToScore[ids[i]] > idToScore[ids[j]]
		}
		return ids[i] < ids[j]
	})
	return ids[0:k]
}
