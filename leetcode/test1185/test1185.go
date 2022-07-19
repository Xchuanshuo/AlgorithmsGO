package test1185

var week = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
var months = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30}

func dayOfTheWeek(day int, month int, year int) string {
	days := day
	days += (year - 1971)*365 + (year-1969)/4
	for _, m := range months[:month-1] {
		days += m
	}
	if month > 2 && (year%400 == 0 || (year%4==0 && year%100!=0)) {
		days++
	}
	return week[(days+3)%7]
}