package main

import "fmt"

func main() {
	var str string
	fmt.Scanln(&str)
	var m = make([]int, 26)
	sum, res := 0, 0
	for _, ch := range str {
		i := int(ch - 'A')
		sum ^= 1 << i
		if (sum>>i)&1 != 0 {
			m[i] = sum
		} else {
			tmp := m[i] ^ sum
			for j := 0; j < 26; j++ {
				if (tmp>>j)&1 != 0 {
					res++
				}
			}
		}
	}
	fmt.Println((res - 26) / 2)
}
