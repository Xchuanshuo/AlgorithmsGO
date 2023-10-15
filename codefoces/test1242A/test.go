package main

/**
 * @Description https://codeforces.com/problemset/problem/1242/A
 * idea: 构造 + gcd.
		1.若n为质数，不存在任何因子，那么答案为n
		2.n不为质数，设n=p*q, 对于所有因子p,q，求gcd即为答案；
         含义为所有间隔距离p、q都可以使用与间隔距离为gcd的格子同一种颜色，所以对于[1,n]
		的所有数，都可以与[1,gcd]中的某个格子使用同一种颜色
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
	var n int64
	fmt.Fscanln(in, &n)
	var res = cal(n)
	fmt.Fprintln(out, res)
	out.Flush()
}

func cal(n int64) int64 {
	var g = n
	for i := int64(2); i*i <= n; i++ {
		if n%i != 0 {
			continue
		}
		g = gcd(g, i)
		g = gcd(g, n/i)
	}
	return g
}

func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
