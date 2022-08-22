package main

import "fmt"

func maxEqualFreq(nums []int) int {
	var count = make(map[int]int)
	var freq = make(map[int]int)
	res, maxFreq := 0, 0
	for i, num := range nums {
		if freq[count[num]] > 0 {
			freq[count[num]]--
		}
		count[num]++
		maxFreq = max(maxFreq, count[num])
		freq[count[num]]++
		isValid := maxFreq == 1 ||
			maxFreq*freq[maxFreq] == i && freq[1] == 1 ||
			maxFreq+(maxFreq-1)*freq[maxFreq-1] == i+1 && freq[maxFreq] == 1
		if isValid {
			res = max(res, i+1)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var nums = []int{2, 2, 1, 1, 5, 3, 3, 5}
	var res = maxEqualFreq(nums)
	fmt.Println(res)
}
