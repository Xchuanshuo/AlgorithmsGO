package lcwt116

var M = int(1e9) + 7

func sumCounts(nums []int) int {
	var n = len(nums)
	var res = 0
	for i := 0; i < n; i++ {
		var set = make(map[int]bool)
		for j := i; j < n; j++ {
			set[nums[j]] = true
			res = (res + len(set)*len(set)) % M
		}
	}
	return res
}

func main() {

}