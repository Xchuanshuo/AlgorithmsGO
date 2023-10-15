package main

/**
 * @Description https://codeforces.com/problemset/problem/1630/A
 * idea: 构造, 特殊元素0(与任何数按位与为0)、n-1(与任何元素按位与结果不变)。对k分情况讨论
		1.k=0, 每个元素与其相对n取反后的元素匹配，值为0(i&(n-1)^i)
		2.0<k<n-1, k与n-1，0与(n-1)^k匹配，其它元素同1, 即 k + 0
		3.k=n-1, n-1与n-2匹配结果为n-2, n-3与1匹配结果为1, 即 n-2 + 1 + 0
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
		var n, k int
		fmt.Fscanf(in, "%d %d", &n, &k)
		fmt.Fscanln(in)
		res, a := cal(n, k)
		if res == -1 {
			fmt.Fprintln(out, -1)
		} else {
			for _, v := range a {
				fmt.Fprintln(out, fmt.Sprintf("%d %d", v[0], v[1]))
			}
		}
		out.Flush()
	}
}

func cal(n int, k int) (int, [][]int) {
	if n == 4 && k == 3 {
		return -1, [][]int{}
	}
	var buildRes = func(mp map[int]bool) [][]int {
		var res = make([][]int, 0)
		for i := 0; i < n; i++ {
			t1, t2 := i, i^(n-1)
			if mp[i] || mp[t2] {
				continue
			}
			res = append(res, []int{t1, t2})
			mp[t1], mp[t2] = true, true
		}
		return res
	}
	var mp = make(map[int]bool)
	if k == 0 {
		return 0, buildRes(mp)
	}
	if k < n-1 {
		mp[k], mp[n-1], mp[0], mp[k^(n-1)] = true, true, true, true
		var res = buildRes(mp)
		res = append(res, []int{k, n - 1})
		res = append(res, []int{0, k ^ (n - 1)})
		return 0, res
	}
	mp[n-2], mp[n-1], mp[n-3] = true, true, true
	mp[0], mp[1], mp[2] = true, true, true
	var res = buildRes(mp)
	res = append(res, []int{n - 2, n - 1})
	res = append(res, []int{n - 3, 1})
	res = append(res, []int{0, 2})
	return 0, res
}
