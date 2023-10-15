package main

/**
 * @Description https://codeforces.com/problemset/problem/1400/C
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
		var s string
		fmt.Fscanln(in, &s)
		var x int
		fmt.Fscanln(in, &x)
		var res = cal(s, x)
		if res == "" {
			fmt.Fprintln(out, -1)
		} else {
			fmt.Fprintln(out, res)
		}
		out.Flush()
	}
}

// 间隔x的0相等，优先处理0再处理1
func cal(s string, x int) string {
	var n = len(s)
	var bs = []byte(strings.Repeat("1", n))
	for idx, v := range s {
		if v == '0' {
			if idx-x >= 0 {
				bs[idx-x] = '0'
			}
			if idx+x < n {
				bs[idx+x] = '0'
			}
		}
	}
	for i := 0; i < n; i++ {
		var v = '0'
		if i-x >= 0 && bs[i-x] == '1' {
			v = '1'
		}
		if i+x < n && bs[i+x] == '1' {
			v = '1'
		}
		if s[i] != uint8(v) {
			return ""
		}
	}
	return string(bs)
}
