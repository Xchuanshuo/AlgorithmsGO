package test2156

var BASE = 1

func subStrHash(s string, power int, modulo int, k int, hashValue int) string {
	BASE = modulo
	n, sum := len(s), 0
	for i := n - 1; i >= n-(k-1); i-- {
		var v = int(s[i] - 'a' + 1)
		sum = power * (sum + v) % modulo
	}
	idx, p := 0, fastPow(power, k-1)
	for i := n - k; i >= 0; i-- {
		var v = int(s[i] - 'a' + 1)
		sum = (sum + v) % modulo
		if sum == hashValue {
			idx = i
		}
		sum = (sum - int(s[i+k-1]-'a'+1)*p%modulo + modulo) % modulo
		sum = sum * power % modulo
	}
	return s[idx : idx+k]
}

func fastPow(a, k int) int {
	if k == 0 {
		return 1
	}
	a %= BASE
	if k%2 == 0 {
		val := fastPow(a, k/2)
		return (val * val) % BASE
	} else {
		return a * fastPow(a, k-1) % BASE
	}
}
