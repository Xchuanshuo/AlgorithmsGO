package test1629

func slowestKey(releaseTimes []int, keysPressed string) byte {
	idx, pre, max := 0, 0, 0
	for i, cur := range releaseTimes {
		if cur - pre > max {
			max = cur - pre
			idx = i
		} else if cur-pre==max && keysPressed[i] > keysPressed[idx] {
			idx = i
		}
		pre = cur
	}
	return keysPressed[idx]
}