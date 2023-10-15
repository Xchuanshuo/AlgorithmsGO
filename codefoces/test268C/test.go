package main

/**
 * @Description https://codeforces.com/problemset/problem/268/C
 * idea: 构造 + 贪心。
		1.点只能取整数且距离为小数，那么在【同一行】或者【同一列】的点距离一定为整数，所以集合中不能出现，
		 且集合最大元素数量为min(n,m)+1
		2.由于每个格子单位为1，对角线距离sqrt(2)肯定不为整数，所以可以考虑取同一条对角线上的所有点，而
		 点不能从(0,0)开始，所以考虑从(0,min(n,m)) ---> (min(n,m), 0)即可
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
	var n, m int
	fmt.Fscanf(in, "%d %d", &n, &m)
	var res = cal(n, m)
	fmt.Fprintln(out, len(res))
	for _, v := range res {
		fmt.Fprintf(out, "%d %d\n", v[0], v[1])
	}
	out.Flush()
}

func cal(n, m int) [][]int {
	var mn = min(n, m)
	var res [][]int
	for i := 0; i <= mn; i++ {
		res = append(res, []int{i, mn - i})
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
