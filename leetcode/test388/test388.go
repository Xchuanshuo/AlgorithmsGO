package main

import (
	"fmt"
)

func lengthLongestPath(input string) int {
	var stack = make([]int, 0)
	var n = len(input)
	var max = 0
	for i := 0; i < n; {
		var depth = 1
		for ; i < n && input[i] == '\t'; i++ {
			depth++
		}
		length, isFile := 0, false
		for ; i < n && input[i] != '\n'; i++ {
			if input[i] == '.' {
				isFile = true
			}
			length++
		}
		i++
		for len(stack) >= depth {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			length += stack[len(stack)-1] + 1
		}
		if isFile {
			if length > max {
				max = length
			}
		} else {
			stack = append(stack, length)
		}
	}
	return max
}

func main() {
	var s = "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"
	var res = lengthLongestPath(s)
	fmt.Println(res)
}
