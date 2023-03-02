package test829

/**
* @Description
* idea: 等差数列求和公式 (a + a+k-1)*k/2 = n ==> (2a+k-1)*k = 2n
       1.2n%k == 0, 2. (2n/k-k+1)%2 == 0
**/

func consecutiveNumbersSum(n int) int {
	var res = 1
	for k := 1; k*k < n; k++ {
		if 2*n%k != 0 {
			continue
		}
		if (2*n/k-(k+1))%2 == 0 {
			res++
		}
	}
	return res
}
