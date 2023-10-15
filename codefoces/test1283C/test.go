package main

/**
 * @Description https://codeforces.com/problemset/problem/1283/C
 * idea: 构造 + 贪心。分情况讨论
		1.如果只剩即没送又没收的，这些人按顺序组成一个环即可
		2.如果只剩没送和没接受的，让每个没送的给没接受的即可
		3.1和2均存在，即没送又没收的组成一个队列，只没送的送给队列第一个人, 队列最后一个人给只没收的
		  剩下情况参考2即可
		理由: 题目保证了每个点的入度一定与出度相等。所以按度分别处理即可
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
	var a = make([]int, 0)
	for i := 0; i < n; i++ {
		var v int
		fmt.Fscanf(in, "%d ", &v)
		a = append(a, v)
	}
	var res = cal(a)
	for _, v := range res {
		fmt.Fprintf(out, "%d ", v)
	}
	fmt.Fprintln(out)
	out.Flush()
}

func cal(fd []int) []int {
	var n = len(fd)
	var inDegree = make([]int, n)
	var outDegree = make([]int, n)
	for i, f := range fd {
		if f != 0 {
			outDegree[i]++
			inDegree[f-1]++
		}
	}
	var waitGive = make([]int, 0)
	var waitRec = make([]int, 0)
	var wait = make([]int, 0)
	for i := 0; i < n; i++ {
		var p = i + 1
		if inDegree[i] == 0 && outDegree[i] == 0 {
			wait = append(wait, p)
		} else if inDegree[i] == 0 {
			waitRec = append(waitRec, p)
		} else if outDegree[i] == 0 {
			waitGive = append(waitGive, p)
		}
	}
	for i := 0; i < len(wait)-1; i++ { // 即没送又接受礼物
		fd[wait[i]-1] = wait[i+1]
	}
	if len(wait) > 0 && len(waitGive) == 0 && len(waitRec) == 0 {
		fd[wait[len(wait)-1]-1] = wait[0]
		return fd
	}
	if len(wait) > 0 { // 第一个待送礼物->没接和没送->第一个待收
		fd[waitGive[0]-1] = wait[0]
		waitGive = waitGive[1:]
		fd[wait[len(wait)-1]-1] = waitRec[0]
		waitRec = waitRec[1:]
	}
	for i, v := range waitGive {
		var w = waitRec[i]
		fd[v-1] = w
	}
	return fd
}
