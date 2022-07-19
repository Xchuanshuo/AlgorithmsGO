package test553

import (
	"fmt"
	"strconv"
	"strings"
)

func optimalDivision(nums []int) string {
	var n = len(nums)
	if n == 1 {
		return strconv.Itoa(nums[0])
	}
	if n == 2 {
		return fmt.Sprintf("%d/%d", nums[0], nums[1])
	}
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%d/(", nums[0]))
	for i := 1; i < n; i++ {
		sb.WriteString(strconv.Itoa(nums[i]))
		if i == n-1 {
			break
		}
		sb.WriteString("/")
	}
	sb.WriteString(")")
	return sb.String()
}
