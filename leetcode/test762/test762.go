package test762

func countPrimeSetBits(left int, right int) int {
	var prims  = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31}
	var primMap = make(map[int]bool)
	for _, v := range prims {
		primMap[v] = true
	}
	var res = 0
	for i := left;i <= right;i++ {
		var cnt = count(i)
		if primMap[cnt] {
			res++
		}
	}
	return res
}

func count(a int) int {
	var res = 0
	for a != 0 {
		a &= a - 1
		res++
	}
	return res
}