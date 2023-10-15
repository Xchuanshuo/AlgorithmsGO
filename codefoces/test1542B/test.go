package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/**
 * @Description https://codeforces.com/problemset/problem/1542/B
 * idea: 构造。生成的任何数都可以写成 a^k+y*b, 所以只需要判断当前(n-a^k)%b == 0 即可
 **/

func main() {
	run(os.Stdin, os.Stdout)
}

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)

	solve := func(curCase int) (_res bool) {
		var n, a, b int
		fmt.Fscanln(in, &n, &a, &b)
		_res = cal(n, a, b)
		if _res {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
		out.Flush()
		return
	}

	cases := 1
	fmt.Fscanln(in, &cases)
	for curCase := 1; curCase <= cases; curCase++ {
		ans := solve(curCase)
		_ = ans
	}

	leftData, _ := io.ReadAll(in)
	if s := strings.TrimSpace(string(leftData)); s != "" {
		panic("有未读入的数据：\n" + s)
	}

}

var cal = func(n, a, b int) bool {
	if a == 1 {
		return (n-1)%b == 0
	}
	var c = 1
	for c <= n {
		if (n-c)%b == 0 {
			return true
		}
		c *= a
	}
	return false
}
