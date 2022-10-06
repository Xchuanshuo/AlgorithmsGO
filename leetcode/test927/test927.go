package test927

/**
 * @Description https://leetcode.cn/problems/three-equal-parts/
 * idea: 1.存储所有值为1的索引位置 2.长度不为3的倍数则不合题意
		 3.数组全部为0的情况下，直接把首尾作为切分点即可
		 4.分为三等分，从每部分1的索引位置开始依次判断是否值相等
 **/

func threeEqualParts(arr []int) []int {
	var ones = make([]int, 0)
	for i, v := range arr {
		if v == 1 {
			ones = append(ones, i)
		}
	}
	if len(ones)%3 != 0 {
		return []int{-1, -1}
	}
	if len(ones) == 0 {
		return []int{0, len(arr) - 1}
	}
	var amount = len(ones) / 3
	i1, i2, i3 := ones[0], ones[amount], ones[amount*2]
	for i3 < len(arr) {
		v1, v2, v3 := arr[i1], arr[i2], arr[i3]
		if v1 != v2 || v2 != v3 || v1 != v3 {
			return []int{-1, -1}
		}
		i1, i2, i3 = i1+1, i2+1, i3+1
	}
	return []int{i1 - 1, i2 - 1}
}
