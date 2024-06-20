package main

func main() {
}

func solve(board [][]byte) {
	if len(board) < 2 || len(board[0]) < 2 {
		return
	}

	// mark cells connected to boarder as boarder_connected('B')
	for i := 0; i < len(board); i++ {
		colorUsingDFS(board, i, 0)               // traverse on first column
		colorUsingDFS(board, i, len(board[i])-1) // traverse on last column
	}

	for j := 0; j < len(board[0]); j++ {
		colorUsingDFS(board, 0, j)            // traverse on first row
		colorUsingDFS(board, len(board)-1, j) // traverse on last row
	}

	// iterate over all cells and mark all 'B' as 'O' and rest as 'X'
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'B' {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}
}

func colorUsingDFS(board [][]byte, i, j int) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[i]) ||
		board[i][j] != 'O' { // if current boarder connected element is not 'O', skip it
		return
	}

	// make current boarder connected element as 'B'
	board[i][j] = 'B'

	// check top cell
	if i-1 >= 0 {
		colorUsingDFS(board, i-1, j)
	}

	// check right cell
	if j+1 < len(board[i]) {
		colorUsingDFS(board, i, j+1)
	}

	// check bottom cell
	if i+1 < len(board) {
		colorUsingDFS(board, i+1, j)
	}

	// check left cell
	if j-1 >= 0 {
		colorUsingDFS(board, i, j-1)
	}
}
