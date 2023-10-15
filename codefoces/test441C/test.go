package main

/**
 * @Description https://codeforces.com/problemset/problem/441/C
 * idea: 构造 + 贪心. 前面k-1个格子每个占两个位置，最后一个占所有格子
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
	var n, m, k int
	fmt.Fscanf(in, "%d %d %d\n", &n, &m, &k)
	var res = cal(n, m, k)
	for _, a := range res {
		fmt.Fprint(out, len(a))
		fmt.Fprint(out, " ")
		for _, v := range a {
			fmt.Fprintf(out, "%d %d ", v[0], v[1])
		}
		fmt.Fprintln(out)
	}
	out.Flush()
}

func cal(n, m, k int) [][][]int {
	var res = make([][][]int, 0)
	var ta = make([][]int, 0)
	var even = true
	var remain = 0
	var t = k
	for i := 0; i < n; i++ {
		even = i%2 == 0
		for p := 0; p < m; p++ {
			var j = p
			if !even {
				j = m - p - 1
			}
			if t > 0 && remain == 0 {
				if t > 1 {
					remain = 2
				} else {
					remain = m*n - (k-1)*2
				}
				if len(ta) > 0 {
					res = append(res, ta)
				}
				ta = make([][]int, 0)
				t--
			}
			ta = append(ta, []int{i + 1, j + 1})
			remain--
		}
	}
	if len(ta) > 0 {
		res = append(res, ta)
	}
	return res
}
