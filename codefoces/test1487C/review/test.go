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
	cases := 1
	fmt.Fscanln(in, &cases)
	for curCase := 1; curCase <= cases; curCase++ {
		var n int
		fmt.Fscanln(in, &n)
		var res = cal(n)
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				fmt.Fprint(out, res[i][j])
				fmt.Fprint(out, " ")
			}
		}
		fmt.Fprintln(out)
		out.Flush()
	}
}

func cal(n int) [][]int {
	var res = make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	var total = n * (n - 1) / 2
	var winCnt = total / n
	for i := 0; i < n; i++ {
		for idx := i + 1; idx < i+1+winCnt; idx++ {
			var j = idx % n
			res[i][j] = 1
			res[j][i] = -1
		}
	}
	return res
}
