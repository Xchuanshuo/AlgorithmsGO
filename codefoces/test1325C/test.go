package main

/**
 * @Description https://codeforces.com/problemset/problem/1325/C
 * idea: 构造 + 贪心。MEX(u,v)为点对u->v的路径中边未出现的【最小自然数】，要使所有点对
        MEX(u,v)的最大值最小，所以分配边权使需要让小边权尽量均匀到每个路径，分情况讨论
		1.根据每个节点的度，若所有节点度均<3, 说明路径是一条链表，不管如何分配，答案还是为n-1
		2.存在度为3的节点，考虑将0，1，2分别分配给节点的3条边，其它权值任意分配即可，这样能保证最小权值一定不超过2
		理由: 树中一个节点到另一个节点的简单路径【最多】只会经过一个节点的两边
		所以除了直接找度为3的节点，同理还可以将0、1、2分配给3个不同的叶子节点
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
	var n int
	fmt.Fscanln(in, &n)
	var a = make([][]int, 0)
	for i := 0; i < n-1; i++ {
		var v, w int
		fmt.Fscanf(in, "%d %d\n", &v, &w)
		a = append(a, []int{v, w})
	}
	var res = cal(a)
	for _, v := range res {
		fmt.Fprintln(out, v)
	}
	out.Flush()
}

func cal(edges [][]int) []int {
	var n = len(edges)
	var res = make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = -1
	}
	var g = make(map[int][]int)
	for i, e := range edges {
		g[e[0]] = append(g[e[0]], i)
		g[e[1]] = append(g[e[1]], i)
	}
	var hasD3 = false
	for i := 1; i <= n+1; i++ {
		if len(g[i]) >= 3 {
			res[g[i][0]] = 0
			res[g[i][1]] = 1
			res[g[i][2]] = 2
			hasD3 = true
			break
		}
	}
	var id = 0
	if hasD3 {
		id = 3
	}
	for i := 0; i < n; i++ {
		if res[i] == -1 {
			res[i] = id
			id += 1
		}
	}
	return res
}
