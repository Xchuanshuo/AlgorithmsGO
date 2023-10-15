package main

/**
 * @Description https://codeforces.com/problemset/problem/1399/D
 * idea: 构造 + 贪心。使用两个栈分别保存前面以0、1结尾的序列，当前为0则去1节尾的序列找，1则去0结尾的序列找
 **/

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	run(os.Stdin, os.Stdout)
}

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	cases := 1
	fmt.Fscanln(in, &cases)
	for curCase := 1; curCase <= cases; curCase++ {
		var n int
		fmt.Fscanln(in, &n)
		var s string
		fmt.Fscanln(in, &s)
		k, res := cal(s)
		fmt.Fprintln(out, k)
		for _, c := range res {
			fmt.Fprint(out, c)
			fmt.Fprint(out, " ")
		}
		fmt.Fprintln(out)
		out.Flush()
	}
}

func cal(s string) (int, []int) {
	var st0 = make([]int, 0)
	var st1 = make([]int, 0)
	var res = make([]int, len(s))
	for i, v := range s {
		var cur = len(st0) + len(st1)
		if v == '0' {
			if len(st1) > 0 {
				cur = st1[len(st1)-1]
				st1 = st1[0 : len(st1)-1]
			}
			st0 = append(st0, cur)
		} else {
			if len(st0) > 0 {
				cur = st0[len(st0)-1]
				st0 = st0[0 : len(st0)-1]
			}
			st1 = append(st1, cur)
		}
		res[i] = cur + 1
	}
	return len(st0) + len(st1), res
}
