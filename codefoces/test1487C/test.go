package main

/**
 * @Description https://codeforces.com/problemset/problem/1487/C
 * idea: 构造 + 贪心。计算总比赛场次total，分情况讨论
		1.若total%n==0，说明每个团队可以输赢得同样的场次，使得分相等
		2.若total%n!=0, 说明每个团队可以赢total/n次，剩下total%n为平局数量，两个一组之间平局即可
		构造方法: 赢: 将1-n个人按顺序围成一圈，每个团队与其后面的total/n个团队比赛取得胜利，
				平局: 每个团队与第n/2个团队比赛获得平局
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
		var res = cal(n)
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				fmt.Fprint(out, res[i][j])
				fmt.Fprint(out, " ")
			}
		}
		fmt.Fprintln(out)
		out.Flush()
	}
}

func cal(n int) [][]int {
	var total = n * (n - 1) / 2
	var winCnt = total / n
	var t = make([][]int, n)
	for i := 0; i < n; i++ {
		t[i] = make([]int, n)
		for j := 0; j < n; j++ {
			t[i][j] = -2
		}
	}
	if total%n != 0 {
		for i := 0; i < n/2; i++ {
			t[i][i+n/2], t[i+n/2][i] = 0, 0
		}
	}
	for i := 0; i < n; i++ {
		cur, j := winCnt, i
		for cur > 0 {
			cur--
			j = (j + 1) % n
			t[i][j], t[j][i] = 1, -1
		}
	}
	return t
}
