package test748

import (
	"strings"
)

func shortestCompletingWord(licensePlate string, words []string) string {
	cnts := make([]int, 128)
	licensePlate = strings.ToLower(licensePlate)
	for i := 0;i < len(licensePlate);i++ {
		c := licensePlate[i]
		if c >= 'a' && c <= 'z' {
			cnts[c]++
		}
	}
	res := ""
	for _, w := range words {
		tCnts := make([]int, 128)
		copy(tCnts, cnts)
		for _, c := range w {
			tCnts[c]--
		}
		isValid := true
		for _, v := range tCnts {
			if v > 0{
				isValid = false
				break
			}
		}
		if isValid && (res == "" || len(w) < len(res)){
			res = w
		}
	}
	return res
}

//func main() {
//	licensePlate := "1s3 PSt"
//	words := []string{"step", "steps", "stripe", "stepple"}
//	res := shortestCompletingWord(licensePlate, words)
//	fmt.Println(res)
//}