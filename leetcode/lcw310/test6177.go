package main

func partitionString(s string) int {
	var cnt = make([]int, 26)
	var res = 0
	for _, v := range s {
		cnt[v-'a']++
		if cnt[v-'a'] > 1 {
			res++
			cnt = make([]int, 26)
			cnt[v-'a'] = 1
		}
	}
	return res + 1
}

//func main() {
//	res := partitionString("ssssss")
//	fmt.Println(res)
//}
