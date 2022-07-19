package main

import (
	"fmt"
	"unicode"
)

var scope map[string][]int
var idx = 0
var exp = ""

func evaluate(expression string) int {
	exp = expression
	scope = make(map[string][]int)
	idx = 0
	return parseExp(expression)
}

func parseExp(str string) int {
	if str[idx] != '(' {
		if unicode.IsLower(rune(str[idx])) {
			var v = parseVar()
			return scope[v][len(scope[v])-1]
		} else {
			return parseInt()
		}
	}
	idx++
	var res = 0
	if str[idx] == 'l' {
		idx += 4
		var vs = make([]string, 0)
		for idx < len(str) {
			if !unicode.IsLower(rune(str[idx])) {
				res = parseExp(str)
				break
			}
			var v = parseVar()
			if str[idx] == ')' {
				res = scope[v][len(scope[v])-1]
				break
			}
			vs = append(vs, v)
			idx++
			var val = parseExp(str)
			if _, exist := scope[v]; !exist {
				scope[v] = make([]int, 0)
			}
			scope[v] = append(scope[v], val)
			idx++
		}
		// 清调当前作用域的栈顶
		for _, v := range vs {
			scope[v] = scope[v][:len(scope[v])-1]
		}
	} else if str[idx] == 'a' {
		idx += 4
		var v1 = parseExp(str)
		idx++
		var v2 = parseExp(str)
		res = v1 + v2
	} else if str[idx] == 'm' {
		idx += 5
		var v1 = parseExp(str)
		idx++
		var v2 = parseExp(str)
		res = v1 * v2
	}
	idx++
	return res
}

func parseVar() string {
	n := len(exp)
	var bytes = make([]byte, 0)
	for idx < n && exp[idx] != ' ' && exp[idx] != ')' {
		bytes = append(bytes, exp[idx])
		idx++
	}
	return string(bytes)
}

func parseInt() int {
	n, val := len(exp), 0
	var neg = false
	if exp[idx] == '-' {
		neg = true
		idx++
	}
	for idx < n && exp[idx] != ' ' && exp[idx] != ')' {
		val = val*10 + int(exp[idx]-'0')
		idx++
	}
	if neg {
		return -val
	}
	return val
}

func main() {
	var exp = "(let x 7 -12)"
	var val = evaluate(exp)
	fmt.Println(val)
}
