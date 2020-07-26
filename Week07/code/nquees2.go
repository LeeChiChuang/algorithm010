package backtrack

var queensAns [][]string

func solveNQueens2(n int) [][]string {
	queensAns = [][]string{}
	board := make([][]byte, n)
	for i:=0; i<n; i++ {
		tmp := make([]byte, n)
		for j:=0; j<n; j++ {
			tmp[j] = '.'
		}
		board[i] = tmp
	}
	queensBack(board, n, 0)

	return queensAns
}

func queensBack(board [][]byte, n, row int) {
	if row == n {
		ans := make([]string, n)
		for i, val := range board {
			ans[i] = string(val)
		}
		queensAns = append(queensAns, ans)
		return
	}

	for col := 0; col < n; col++ {
		if !isValidQueens(board, row, col) {
			continue
		}
		board[row][col] = 'Q'
		queensBack(board, n, row+1)

		// 撤销选择
		board[row][col] = '.'
	}
}

func isValidQueens(board [][]byte, row, col int) bool {
	n := len(board)

	for i:=0; i<n; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}

	// '/'
	for i, j := row-1, col+1; i>=0 && j<n; i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	// '\'
	for i,j := row-1, col-1; i>=0 && j>=0; i,j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	return true
}
