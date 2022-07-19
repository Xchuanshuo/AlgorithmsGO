package test419

func countBattleships(board [][]byte) int {
	res := 0
	m, n := len(board), len(board[0])
	for i := 0;i < m;i++ {
		for j := 0; j < n; j++{
			if board[i][j] != 'X' {
				continue
			}
			res++
			if i > 0 && board[i-1][j] == 'X' {
				res--
			}
			if j > 0 && board[i][j-1] == 'X' {
				res--
			}
		}
	}
	return res
}
