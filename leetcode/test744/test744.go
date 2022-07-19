package test744

func nextGreatestLetter(letters []byte, target byte) byte {
	var n = len(letters)
	l, r := 0, n-1
	for l <= r {
		mid := l + (r-l)/2
		if letters[mid] <= target {
			l = mid + 1
		} else {
			if mid == 0 || letters[mid-1] <= target {
				return letters[mid]
			} else {
				r = mid - 1
			}
		}
	}
	return letters[0]
}
