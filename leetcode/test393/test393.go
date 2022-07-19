package main

import "fmt"

const (
	MAGIC   = 0b10000000
	SECOND  = 0b11000000
	THIRD   = 0b11100000
	FOUR    = 0b11110000
	INVALID = 0b11111000
)

func validUtf8(data []int) bool {
	var cnt = 0
	for _, d := range data {
		if cnt == 0 {
			if d&INVALID == INVALID {
				return false
			}
			if d&FOUR == FOUR {
				cnt = 3
			} else if d&THIRD == THIRD {
				cnt = 2
			} else if d&SECOND == SECOND {
				cnt = 1
			} else if d>>7 == 1 {
				return false
			}
		} else {
			if cnt < 0 || d>>6 != 0b10 {
				return false
			}
			cnt -= 1
		}
	}
	return cnt == 0
}

func main() {
	var arr = []int{197, 130, 1}
	res := validUtf8(arr)
	fmt.Println(res)
}
