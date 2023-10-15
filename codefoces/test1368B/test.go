package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/**
 * @Description https://codeforces.com/problemset/problem/1542/B
 * idea: 构造。贪心, 均匀添加目标串中的字符即可
 **/

func main() {
	run(os.Stdin, os.Stdout)
}

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	cases := int64(0)
	fmt.Fscanln(in, &cases)
	fmt.Fprintln(out, cal(cases))
}

var t = "codeforces"

func cal(k int64) string {
	var cnt = make([]int, len(t))
	for i := 0; i < len(t); i++ {
		cnt[i] = 1
	}
	if k == 1 {
		return t
	}
	var idx = 0
	for true {
		var total = int64(1)
		cnt[idx]++
		for i := 0; i < len(t); i++ {
			total *= int64(cnt[i])
		}
		if total >= k {
			break
		}
		idx = (idx + 1) % len(t)
	}
	var bs = make([]byte, 0)
	for i := 0; i < len(t); i++ {
		for j := 0; j < cnt[i]; j++ {
			bs = append(bs, t[i])
		}
	}
	return string(bs)
}
