package test1518

func numWaterBottles(numBottles int, numExchange int) int {
	res := numBottles
	for numBottles >= numExchange {
		res += numBottles / numExchange
		numBottles = numBottles / numExchange + numBottles%numExchange
	}
	return res
}