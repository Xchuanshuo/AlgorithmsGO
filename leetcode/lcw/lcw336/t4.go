package lcw336

import "sort"

func findMinimumTime(tasks [][]int) int {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i][1] < tasks[j][1]
	})
	var visited = make(map[int]bool)
	var res = 0
	for _, ta := range tasks {
		s, e, t := ta[0], ta[1], ta[2]
		for i := s; i <= e; i++ {
			if visited[i] {
				t--
			}
		}
		for i := e; t > 0; i-- {
			if !visited[i] {
				visited[i] = true
				t--
				res++
			}
		}
	}
	return res
}
