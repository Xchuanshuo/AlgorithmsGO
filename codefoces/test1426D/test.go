package main

/**
 * @Description https://codeforces.com/problemset/problem/1426/D
 * idea: 构造 + 贪心。
		1.找和为0的字段 => 同两数之和, hash表是否存在相等的前缀和
		2.找到一个区间后，贪心.在当前元素之前添加一个无穷大的元素，使得到当前位置的前缀和>后面的，
		 此时答案+1，将hash表清空，重新重复前面的步骤
	   此种办法最优, 只用了一次操作就删除了当前位置结尾的所有和为0的区间
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
	var a = make([]int, 0)
	for curCase := 1; curCase <= cases; curCase++ {
		var v int
		fmt.Fscanf(in, "%d ", &v)
		a = append(a, v)
	}
	var res = cal(a)
	fmt.Fprintln(out, res)
	out.Flush()
}

func cal(a []int) int {
	var res = 0
	var sum int64
	var set = map[int64]bool{0: true}
	for _, v := range a {
		sum += int64(v)
		if set[sum] {
			res++
			set = map[int64]bool{0: true}
			sum = int64(v)
		}
		set[sum] = true
	}
	return res
}
