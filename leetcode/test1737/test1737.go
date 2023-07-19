package test1737

/**
 * @Description https://leetcode.cn/problems/change-minimum-characters-to-satisfy-one-of-three-conditions/
 * idea: 1.全改成同一个字母 枚举所有字母
		2.枚举边界字母，使得所有A <= 边界，所有B > 边界
					     所有A > 边界，所有B <= 边界
 **/

func minCharacters(a string, b string) int {
	cnt1, cnt2 := make([]int, 26), make([]int, 26)
	for _, v := range a {
		cnt1[int(v-'a')]++
	}
	for _, v := range b {
		cnt2[int(v-'a')]++
	}
	var res = len(a) + len(b)
	for i := 0; i < 26; i++ {
		var cur = len(a) - cnt1[i] + len(b) - cnt2[i]
		res = min(res, cur)
	}
	curA, curB := 0, 0
	for i := 0; i < 25; i++ {
		curA += cnt1[i]
		curB += cnt2[i]
		res = min(res, len(a)-curA+curB)
		res = min(res, curA+len(b)-curB)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
