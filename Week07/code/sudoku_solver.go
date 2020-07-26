package backtrack

func solveSudoKu(board [][]byte) {
	sudoBackTrack(board, 0, 0)
}

func sudoBackTrack(board [][]byte, row, col int) bool {
	// 列走完 进入下一行
	if col == len(board) {
		return sudoBackTrack(board, row+1, 0)
	}
	// 行走完了 成功
	if row == len(board) {
		return true
	}
	// 是空格下一列
	if board[row][col] != '.' {
		return sudoBackTrack(board, row, col+1)
	}

	// 选择1-9
	for i := byte('1'); i <= '9'; i++ {
		if ! isValidSudo(board, i,  row, col) {
			continue
		}
		board[row][col] = i
		if sudoBackTrack(board, row, col+1) {
			return true
		}
		board[row][col] = '.'
	}

	return false
}

func isValidSudo(board [][]byte, num byte, row, col int) bool {
	for i:=0; i<len(board); i++ {
		if board[row][i] == byte(num) {
			return false
		}
		if board[i][col] == byte(num) {
			return false
		}
		if board[(row/3)*3 + i/3][(col/3)*3+i%3] == byte(num) {
			return  false
		}
	}
	return  true
}
