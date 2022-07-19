package main

import (
	"fmt"
	"strings"
)

func solve(str1 string, str2 string) {
	tMap := make(map[int]bool)
	for idx, v := range str1 {
		c := (v-'0')^1
		var sb strings.Builder
		sb.WriteString(str1[0:idx])
		sb.WriteRune('0' + c)
		sb.WriteString(str1[idx+1:])
		val := calTwo(sb.String())
		tMap[val] = true
	}
	found := false
	res := 0
	for idx, v := range str2 {
		cs := make([]int, 0)
		if v == '0' {
			cs = append(cs, 1, 2)
		} else if v == '1' {
			cs = append(cs, 0, 2)
		} else if v == '2' {
			cs = append(cs, 0, 1)
		}
		for _, c := range cs {
			var sb strings.Builder
			sb.WriteString(str2[0:idx])
			sb.WriteRune(rune(c + '0'))
			sb.WriteString(str2[idx+1:])
			val := calThree(sb.String())
			if tMap[val] {
				found = true
				res = val
				break
			}
		}
		if found { break }
	}
	fmt.Println(res)
}

func calTwo(s string) int {
	res := 0
	for _, v := range s {
		res = res * 2 + int(v-'0')
	}
	return res
}

func calThree(s string) int {
	res := 0
	for _, v := range s {
		res = res * 3 + int(v-'0')
	}
	return res
}


func main() {
	//input := bufio.NewScanner(os.Stdin)
	//input.Scan()
	//str1 := input.Text()
	//input.Scan()
	//str2 := input.Text()
	//solve(str1, str2)
	var a int
	var s string
	fmt.Scanln(&a)
	fmt.Scanln(&s)
	var arr = make([]int, 10)
	for i := 0;i < 10;i++ {
		fmt.Scanln(&arr[i])
	}
	println(a)
	println(s)
	fmt.Print(arr)
}


