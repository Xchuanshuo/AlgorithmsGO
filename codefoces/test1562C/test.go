package main

/**
 * @Description https://codeforces.com/problemset/problem/1562/C
 * idea: 构造 + 贪心。考虑k取尽可能小的数，如0、1、2，取0时需区间1、2都为0；取1时，考虑在区间1前
							面加一个0，值不变；取2时，考虑在区间1后面添加一个0。
		1.首先考虑字符串不存在0，那么区间1、2选个数相同的1即可, [1,n-1], [2,n]
		2.存在0。若0在左半边，位置为p，此时可以选[p,n], [p+1,n]
				若0在右半边，位置为p，此时可以选[1,p], [1,p-1]
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
	cases := 1
	fmt.Fscanln(in, &cases)
	for curCase := 1; curCase <= cases; curCase++ {
		var n int
		fmt.Fscanln(in, &n)
		var s string
		fmt.Fscanln(in, &s)
		var res = cal(s)
		fmt.Fprintf(out, "%d %d %d %d\n",
			res[0], res[1], res[2], res[3])
		out.Flush()
	}
}

func cal(s string) []int {
	var n = len(s)
	var mid = n / 2
	for i, v := range s {
		var p = i + 1
		if v == '1' {
			continue
		}
		if p <= mid {
			return []int{p, n, p + 1, n}
		} else {
			return []int{1, p, 1, p - 1}
		}
	}
	return []int{1, n - 1, 2, n}
}
