package test593

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	var set = make(map[int]bool)
	set[distance(p1, p2)] = true
	set[distance(p1, p3)] = true
	set[distance(p1, p4)] = true
	set[distance(p2, p3)] = true
	set[distance(p2, p4)] = true
	set[distance(p3, p4)] = true
	if len(set) != 2 || set[0] {
		return false
	}
	var a = make([]int, 0)
	for k := range set {
		a = append(a, k)
	}
	if a[0] > a[1] {
		return a[0] == 2*a[1]
	}
	return a[1] == 2*a[0]
}

func distance(a, b []int) int {
	var v1 = a[0] - b[0]
	var v2 = a[1] - b[1]
	return v1*v1 + v2*v2
}
