package main

/**
 * @Description https://codeforces.com/problemset/problem/1368/C
 * idea: 构造 + 贪心。斜对角为为中心，构造n个点，周围4个点都需要染色
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
	var n int
	fmt.Fscanln(in, &n)
	var res = cal(n)
	fmt.Fprintln(out, len(res))
	for _, v := range res {
		fmt.Fprintf(out, "%d %d\n", v[0], v[1])
	}
	out.Flush()
}

func cal(n int) [][]int {
	var res = make([][]int, 0)
	// 中间
	for i := 0; i <= n+1; i++ {
		res = append(res, []int{i, i})
	}
	// 两侧
	x, y := 0, 1
	for i := 0; i <= n; i++ {
		res = append(res, []int{x + i, y + i})
	}
	x, y = 1, 0
	for i := 0; i <= n; i++ {
		res = append(res, []int{x + i, y + i})
	}
	return res
}
