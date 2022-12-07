package test1106

func parseBoolExpr(expression string) bool {
	var s = make([]int32, 0)
	for _, e := range expression {
		if isOp(e) {
			s = append(s, e)
		} else if e == 't' || e == 'f' {
			s = append(s, e)
		} else if e == ')' {
			var conds = make([]bool, 0)
			for len(s) > 0 && !isOp(s[len(s)-1]) {
				if s[len(s)-1] == 't' {
					conds = append(conds, true)
				} else {
					conds = append(conds, false)
				}
				s = s[0 : len(s)-1]
			}
			var op = s[len(s)-1]
			s = s[0 : len(s)-1]
			s = append(s, calc(op, conds))
		}
	}
	return s[0] == 't'
}

func isOp(b int32) bool {
	if b == '|' || b == '&' || b == '!' {
		return true
	}
	return false
}

func calc(c int32, conditions []bool) int32 {
	switch c {
	case '|':
		for _, cond := range conditions {
			if cond == true {
				return 't'
			}
		}
		return 'f'
	case '&':
		for _, cond := range conditions {
			if cond == false {
				return 'f'
			}
		}
		return 't'
	case '!':
		if conditions[0] {
			return 'f'
		}
		return 't'
	}
	return 't'
}
