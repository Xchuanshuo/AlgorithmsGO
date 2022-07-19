package test1175

const M int64 = 1e9 + 7

func numPrimeArrangements(n int) int {
	pa, pb := int64(1), int64(1)
	cntA, cntB := int64(0), int64(1)
	for i := 2; i <= n; i++ {
		if isPrim(i) {
			cntA++
			pa = (pa * cntA) % M
		} else {
			cntB++
			pb = (pb * cntB) % M
		}
	}
	return int((pa * pb) % M)
}

func isPrim(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
