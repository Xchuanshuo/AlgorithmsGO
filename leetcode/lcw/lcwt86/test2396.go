package lcwt86

func isStrictlyPalindromic(n int) bool {
	for i := 2; i <= n-2; i++ {
		var arr = ToBaseArr(n, i)
		j, length := 0, len(arr)
		for j <= length/2 {
			if arr[j] != arr[length-j-1] {
				return false
			}
			j++
		}
	}
	return true
}

func ToBaseArr(n int, base int) []int {
	var arr = make([]int, 0)
	var t = n
	for t > 0 {
		arr = append(arr, t%base)
		t /= base
	}
	return arr
}
