package main

/**
 * @Description https://codeforces.com/problemset/problem/1809/C
 * idea: 构造 + 贪心。
		1.首先考虑用连续正数去构造，连续t个正数(考虑使用2，因为子数组不能为0，使用1计算不方便)，
		 可形成的非空区间个数为t*(t+1)/2，求最小的t, 满足t*(t+1)/2 <= k， 后续其它位置直接使用-1000
		2.若t个正数构成的区间等于k, 直接返回答案即可
		 若t个正数构成的区间小于k，考虑修改第t+1个数, 补齐剩余的区间个数d = k-t*(t+1)/2
		 且由于t+1比t构成的区间个数多t个，所以肯定通过修改t+1个数补齐区间；只需要让[0,t-1]的
		 后t-d个数之和小于第t个数即可，最终修改的值为 -2*(t-d)-1 (-1是为了防止区间和为0)
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
			fmt.Fprintf(out, "%d ", v)
		}
		fmt.Fprintln(out)
		out.Flush()
	}
}

func cal(n, k int) []int {
	var res = make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = -1000
	}
	var t = 1
	for true {
		if t*(t+1)/2 > k {
			t -= 1
			break
		}
		t++
	}
	for i := 0; i < t; i++ {
		res[i] = 2
	}
	if t*(t+1)/2 == k {
		return res
	}
	var d = k - t*(t+1)/2
	res[t] = -2*(t-d) - 1
	return res
}

// n*(n+1)/2
// 1: 1,
// 2: 3,
// 3: 6,
// 4: 10,
// 5: 15,
// 6: 21

// n=6, k=18, t=5, d=3, 2 2 2 2 2

// abb ba abb
// abb ba aaa
