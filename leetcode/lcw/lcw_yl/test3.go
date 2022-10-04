package main

func storedEnergy(storeLimit int, power []int, supply [][]int) int {
	var res = 0
	curMin, curMax := 0, 0
	for i, v := range power {
		if len(supply) > 0 && i == supply[0][0] {
			curMin, curMax = supply[0][1], supply[0][2]
			supply = supply[1:]
		}
		if v < curMin {
			res = max(0, res-(curMin-v))
		} else if v > curMax {
			res = min(res+v-curMax, storeLimit)
		}
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
