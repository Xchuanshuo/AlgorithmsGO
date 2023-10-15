package main

/**
 * @Description https://codeforces.com/problemset/problem/1352/F
 * idea: 构造 + 贪心。
 **/

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
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
		var n0, n1, n2 int
		fmt.Fscanf(in, "%d %d %d\n", &n0, &n1, &n2)
		var res = cal(n0, n1, n2)
		fmt.Fprintln(out, res)
		out.Flush()
	}
}

// 间隔x的0相等，优先处理0再处理1
func cal(n0, n1, n2 int) string {
	var bs []byte
	if n0 > 0 {
		bs = []byte(strings.Repeat("0", n0+1))
	}
	if n2 > 0 {
		if n0 > 0 {
			n1 -= 1
		}
		bs = append(bs, strings.Repeat("1", n2+1)...)
	}
	if len(bs) == 0 {
		bs = append(bs, '1')
	}
	if n1 > 0 && bs[len(bs)-1] != '1' {
		n1 -= 1
		bs = append(bs, '1')
	}

	var v = 0
	for n1 > 0 {
		bs = append(bs, byte('0'+v))
		n1--
		v ^= 1
	}
	return string(bs)
}
