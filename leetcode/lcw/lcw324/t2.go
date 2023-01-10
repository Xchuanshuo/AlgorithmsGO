package lcw324

var table = make([]int, 0)

func init() {
	var isPrim = func(n int) bool {
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}
	for i := 2; i <= int(1e5); i++ {
		if isPrim(i) {
			table = append(table, i)
		}
	}
}

func smallestValue(n int) int {
	res, t := n, n
	var pre = n
	for t > 1 {
		var sum = 0
		for i := 0; i < len(table); {
			var c = table[i]
			if t == 1 {
				break
			}
			if t%c != 0 {
				i++
				continue
			}
			sum += c
			t /= c
		}
		if t == 1 {
			res = min(res, sum)
		}
		if pre == sum {
			break
		}
		t = sum
		pre = t
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
