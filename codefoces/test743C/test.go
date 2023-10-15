package main

/**
 * @Description https://codeforces.com/problemset/problem/743/C
 * idea: 构造. 2/n拆分成1/n+1/n => 求两个数和为1/n, 任选一1/(n+1)，可以得出另一个数1/n*(n+1)
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
	if n == 1 {
		fmt.Fprintln(out, -1)
	} else {
		fmt.Fprintf(out, "%d %d %d\n", n, n+1, n*(n+1))
	}
	out.Flush()
}
