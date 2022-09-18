package main

import "sort"

func matchPlayersAndTrainers(players []int, trainers []int) int {
	sort.Ints(players)
	sort.Ints(trainers)
	pr, tr := len(players)-1, len(trainers)-1
	var res = 0
	for tr >= 0 && pr >= 0 {
		if trainers[tr] >= players[pr] {
			tr--
			pr--
			res++
		} else {
			pr--
		}
	}
	return res
}
