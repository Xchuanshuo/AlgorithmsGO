package main

import "fmt"

func modifyString(s string) string {
	var chs = []rune(s)
	pre := rune(-1)
	for i := 0;i < len(chs);i++ {
		if chs[i] == '?' {
			next := rune(-1)
			if i+1 < len(chs) { next = chs[i+1] }
			for j := 0;j < 26;j++ {
				var r = rune('a' + j)
				if r != pre && r != next {
					chs[i] = r
					break
				}
			}
		}
		pre = chs[i]
	}
	return string(chs)
}

func main() {
	res := modifyString("a?")
	fmt.Println(res)
}