package main

/**
 * @Description https://codeforces.com/problemset/problem/1800/E2
 * idea: 构造 + 贪心。 1.分组 2.组内排序后比较字符串
 **/

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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
		var s1, s2 string
		fmt.Fscanln(in, &s1)
		fmt.Fscanln(in, &s2)
		var res = cal(s1, s2, k)
		fmt.Fprintln(out, res)
		out.Flush()
	}
}

func cal(s, t string, k int) string {
	var n = len(s)
	var waits = make([]int, 0)
	for i := 0; i < n; i++ {
		if s[i] != t[i] {
			waits = append(waits, i)
		}
	}
	var gs = make(map[int][]int)
	var vis = make([]int, n)
	var dfs func(idx, id int)
	dfs = func(idx, id int) {
		if vis[idx] > 0 {
			return
		}
		vis[idx] = id
		gs[id] = append(gs[id], idx)
		if idx+k < n {
			dfs(idx+k, id)
		}
		if idx+k+1 < n {
			dfs(idx+k+1, id)
		}
		if idx-k >= 0 {
			dfs(idx-k, id)
		}
		if idx-k-1 >= 0 {
			dfs(idx-k-1, id)
		}
	}
	// 1.分组
	var id = 1
	for _, v := range waits {
		if vis[v] == 0 {
			dfs(v, id)
			id++
		}
	}
	// 2.每个分组是否能通过交换得到字符串s2
	for _, g := range gs {
		var bs1 = make([]byte, 0)
		var bs2 = make([]byte, 0)
		for _, v := range g {
			bs1 = append(bs1, s[v])
			bs2 = append(bs2, t[v])
		}
		sort.Slice(bs1, func(i, j int) bool {
			return bs1[i] < bs1[j]
		})
		sort.Slice(bs2, func(i, j int) bool {
			return bs2[i] < bs2[j]
		})
		if string(bs1) != string(bs2) {
			return "NO"
		}
	}
	return "YES"
}