package main

func canFormArray(arr []int, pieces [][]int) bool {
	var idMap = make(map[int]int)
	for idx, a := range arr {
		idMap[a] = idx + 1
	}
	for _, piece := range pieces {
		var last = idMap[piece[0]]
		if last == 0 {
			return false
		}
		for i := 1; i < len(piece); i++ {
			if idMap[piece[i]]-last != 1 {
				return false
			}
			last = idMap[piece[i]]
		}
	}
	return true
}

func main() {
	var a1 = []int{85}
	var a2 = [][]int{{85}}
	canFormArray(a1, a2)
}
