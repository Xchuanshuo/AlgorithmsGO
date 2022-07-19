package main

import "fmt"

func isAdditiveNumber(num string) bool {
	for i := 1;i < len(num);i++ {
		for j := i+1; j < len(num); j++{
			if isValid(num, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func isValid(num string, i, j, k int) bool {
	if num[i] == '0' && j != i+1 { return false }
	if num[j] == '0' && k != j+1 { return false }
	a, b := num[i:j], num[j:k]
	sum := addTwoNum(a, b)
	if k + len(sum) > len(num) { return false }
	for x := 0;x < len(sum);x++ {
		if sum[x] != num[k+x] { return false }
	}
	if k + len(sum) == len(num) { return true }
	return isValid(num, j, k, k + len(sum))
}

func addTwoNum(a string, b string) string {
	i, j := len(a)-1, len(b)-1
	carry := byte(0)
	s := make([]rune, 0)
	for i>=0 || j>= 0 || carry>0{
		var va, vb byte
		if i >= 0 { va = a[i]-'0'}
		if j >= 0 { vb = b[j]-'0'}
		v := va + vb + carry
		s = append(s, rune(v%10)+'0')
		carry = v/10
		i--
		j--
	}
	str := string(s)
	return Reverse(str)
}

func Reverse(s string) string {
	r := []rune(s)
	for i,j := 0, len(s)-1;i < j;i,j = i+1,j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	res := isAdditiveNumber("112358")
	fmt.Println(res)
}