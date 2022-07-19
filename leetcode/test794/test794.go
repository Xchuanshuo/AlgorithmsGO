package test794

func validTicTacToe(board []string) bool {
	cntO, cntX := 0, 0
	for _, s := range board {
		for i := 0;i < len(s);i++ {
			if s[i] == 'O' {
				cntO++
			} else if s[i] == 'X' {
				cntX++
			}
		}
	}
	if cntX != cntO && cntX - 1 != cntO {
		return false
	}
	if isWin(board, 'O') && cntX != cntO {
		return false
	}
	if isWin(board, 'X') && cntX - 1 != cntO {
		return false
	}
	return true
}

func isWin(board []string, c uint8) bool {
	// 判断行列
	for i := 0;i < 3;i++ {
		if board[i][0] == c && board[i][1] == c && board[i][2] == c {
			return true
		}
	}
	for i := 0;i < 3;i++ {
		if board[0][i] == c && board[1][i] == c && board[2][i] == c {
			return true
		}
	}
	// 判断斜对角
	if board[0][0] == c && board[1][1] == c && board[2][2] == c {
		return true
	}
	if board[0][2] == c && board[1][1] == c && board[2][0] == c {
		return true
	}
	return false
}