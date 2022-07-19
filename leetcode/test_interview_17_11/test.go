package test_interview_17_11

func findClosest(words []string, word1 string, word2 string) int {
	start, end := -1, -1
	res := len(words)
	for i, w := range words {
		if w == word1 {
			start = i
		}
		if w == word2 {
			end = i
		}
		if start == -1 || end == -1 {
			continue
		}
		res = min(res, abs(start-end))
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
