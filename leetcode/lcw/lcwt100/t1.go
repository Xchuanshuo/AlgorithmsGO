package lcwt100

func distMoney(money int, children int) int {
	if children > money {
		return -1
	}
	money -= children
	var res = 0
	for i := 0; i < children; i++ {
		if money == 3 && i == children-1 {
			res--
			return res
		}
		if money < 7 {
			return res
		}
		money -= 7
		res++
	}
	if money != 0 {
		res--
	}
	return res
}
