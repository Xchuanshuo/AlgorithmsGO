package main

import (
	"fmt"
	"strconv"
	"unicode"
)

const (
	Init = iota
	None
	Let
	Let1
	Let2
	Add
	Add1
	Add2
	Mul
	Mul1
	Mul2
	Done
)

type Node struct {
	state  int
	val    int
	vr     string
	e1, e2 int
}

func evaluate(exp string) int {
	var n = len(exp)
	var scopes = make(map[string][]int)
	var calToken = func(str string) int {
		if !unicode.IsLower(rune(str[0])) {
			v, _ := strconv.Atoi(str)
			return v
		}
		vals := scopes[str]
		return vals[len(vals)-1]
	}
	var vals = make([][]string, 0)
	var s = make([]Node, 0)
	var cur = Node{state: Init}
	var token = ""
	for i := 0; i < n; {
		if exp[i] == ' ' {
			i++
			continue
		}
		if exp[i] == '(' {
			i++
			s = append(s, cur)
			cur = Node{state: None}
			continue
		}
		if exp[i] == ')' {
			i++
			if cur.state == Let2 {
				token = strconv.Itoa(cur.val)
				// 清除作用域
				for _, v := range vals[len(vals)-1] {
					scopes[v] = scopes[v][0 : len(scopes[v])-1]
				}
				vals = vals[0 : len(vals)-1]
			} else if cur.state == Add2 {
				token = strconv.Itoa(cur.e1 + cur.e2)
			} else if cur.state == Mul2 {
				token = strconv.Itoa(cur.e1 * cur.e2)
			}
			cur, s = s[len(s)-1], s[0:len(s)-1]
		} else {
			var s = i
			for i < n && exp[i] != ' ' && exp[i] != ')' {
				i++
			}
			token = exp[s:i]
		}
		switch cur.state {
		case Init:
			cur.val, _ = strconv.Atoi(token)
			cur.state = Done
		case None:
			if token == "let" { // 创建新的作用域
				cur.state = Let
				vals = append(vals, []string{})
			} else if token == "add" {
				cur.state = Add
			} else if token == "mult" {
				cur.state = Mul
			}
		case Let:
			if exp[i] == ')' {
				cur.val = calToken(token)
				cur.state = Let2
			} else {
				// 记录当前作用域的变量名
				vals[len(vals)-1] = append(vals[len(vals)-1], token)
				cur.vr = token
				cur.state = Let1
			}
		case Let1:
			if _, exist := scopes[cur.vr]; !exist {
				scopes[cur.vr] = make([]int, 0)
			}
			scopes[cur.vr] = append(scopes[cur.vr], calToken(token))
			cur.state = Let
		case Add:
			cur.state = Add1
			cur.e1 = calToken(token)
		case Add1:
			cur.state = Add2
			cur.e2 = calToken(token)
		case Mul:
			cur.state = Mul1
			cur.e1 = calToken(token)
		case Mul1:
			cur.state = Mul2
			cur.e2 = calToken(token)
		}
	}
	return cur.val
}

func main() {
	//var exp = "(let x 7 -12)"
	var exp = "(let x 2 (mult x (let x 3 y 4 (add x y))))"
	var val = evaluate(exp)
	fmt.Println(val)
}
