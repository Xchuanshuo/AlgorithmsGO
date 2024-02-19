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
		var n, k int
		fmt.Fscanf(in, "%d %d\n", &n, &k)
		var res = cal(n, k)
		for _, v := range res {
			fmt.Fprint(out, v)
			fmt.Fprint(out, " ")
		}
		fmt.Fprintln(out)
		out.Flush()
	}
}

func cal(n, k int) []int {
	var res = make([]int, 0)
	var m = n - k
	for i := 1; i <= k-m-1; i++ {
		res = append(res, i)
	}
	for i := k; i >= k-m; i-- {
		res = append(res, i)
	}
	return res
}
