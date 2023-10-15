package main

/**
 1 2 3 4 3 2 1
 * @Description https://codeforces.com/problemset/problem/1473/C
 * idea: 构造 + 贪心。
		重要结论: 对于一个2*p长度的序列s[1],s[2]...s[p], s[p-1], s[p-2]...s[1]，
				不管s为升序还是降序，逆序对数量都不变，为(p-1)^2
		证明: 任取两个不相同的数字x和y, 在该序列位置为 x..y..y..x或者x..y..x，无论x和y谁大，前者逆序对均为2，后者均为1
			 所以总的逆序对数目为 2*(p-1)*(p-2)/2 + (p-1) = (p-1)^2，即逆序对数量为以位置p为对称中心的2*(1,p-1)和以p为逆序对左边数的p-1个
		设 m = n-k
		1.由于b序列逆序对数量不能超过a，所以前面1..k-m-1保持不变(1..k-m-1递增)，对逆序对不产生贡献
		2.字典序尽可能大，且逆序对数量与正序还是逆序无关，所以后k-m..k个数倒序即可
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
