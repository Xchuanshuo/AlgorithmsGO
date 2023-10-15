package main

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
	fmt.Fscanf(in, "%d %d\n", &n, &m)
	var res = cal(n, m)
	fmt.Fprintln(out, len(res))
	for _, v := range res {
		fmt.Fprintf(out, "%d %d\n", v[0], v[1])
	}
	out.Flush()
}

func cal(n, m int) [][]int {
	var res = make([][]int, 0)
	var k = min(n, m)
	for i := 0; i <= k; i++ {
		res = append(res, []int{i, k - i})
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
