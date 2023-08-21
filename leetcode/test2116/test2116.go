package test2116

/**
 * @Description https://leetcode.cn/problems/check-if-a-parentheses-string-can-be-valid/
 * idea: 可变字符先保存起来，关键点: 考虑当前不可变字符什么情况【无法】被后续字符【消除】
		1.从左到右遍历，当前右括号数量大于左边的左括号+可变字符
		2.从右到左遍历，当前左括号数量大于右边的右括号+可变字符
 **/

func canBeValid(s string, locked string) bool {
	var n = len(s)
	if n%2 != 0 {
		return false
	}
	a, b, c := 0, 0, 0
	for i, v := range s {
		if locked[i] == '0' {
			c++
		} else {
			if v == ')' {
				a++
			} else {
				b++
			}
		}
		if a > b+c {
			return false
		}
	}
	a, b, c = 0, 0, 0
	for i := n - 1; i >= 0; i-- {
		if locked[i] == '0' {
			c++
		} else {
			if s[i] == '(' {
				a++
			} else {
				b++
			}
		}
		if a > b+c {
			return false
		}
	}
	return true
}
