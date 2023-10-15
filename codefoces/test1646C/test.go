package main

/**
 * @Description https://codeforces.com/problemset/problem/1646/C
 * idea: 构造 + 贪心。
		1.枚举所有阶乘组合
		2.n-阶乘后，二进制1的个数与阶乘数的个数即为k
 **/

import (
	"bufio"
	"fmt"
	"io"
	"math/bits"
	"os"
)

var facs []int64
var INF = int64(1e12)

func init() {
	var p = int64(1)
	var total = int64(1)
	for total < INF {
		facs = append(facs, total)
		p += 1
		total *= p
	}
}

func main() {
	run(os.Stdin, os.Stdout)
}

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	cases := 1
	fmt.Fscanln(in, &cases)
	for curCase := 1; curCase <= cases; curCase++ {
		var n int64
		fmt.Fscanln(in, &n)
		var res = cal(n)
		fmt.Fprintln(out, res)
		out.Flush()
	}
}

func cal(n int64) int {
	var m = len(facs)
	var l = 1 << m
	var res = bits.OnesCount64(uint64(n))
	for i := 0; i < l; i++ {
		var sum = int64(0)
		var cnt = 0
		var selected = make(map[int64]bool)
		for j := 0; j < m; j++ {
			if i&(1<<j) != 0 {
				sum += facs[j]
				selected[facs[j]] = true
				cnt++
			}
		}
		var cur = n - sum
		if cur < 0 {
			continue
		}
		var valid = true
		for j := 0; j < 60; j++ { // 校验所选阶乘数与二进制数是否有冲突
			if cur&(1<<j) != 0 && selected[1<<j] {
				valid = false
				break
			}
		}
		if !valid {
			continue
		}
		res = min(res, cnt+bits.OnesCount64(uint64(cur)))
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
