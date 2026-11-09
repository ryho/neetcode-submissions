
var allBoards [][]string
func solveNQueens(n int) [][]string {
	allBoards = nil
	board := emptyBoard(n)
	solveNQueensHelper(board, 0)
	return allBoards
}

func solveNQueensHelper(board [][]byte, i int) {
	for j := range board {
		if canPlaceQueen(board, i, j) {
			board[i][j] = 'Q'
			if i == len(board)-1 {
				allBoards = append(allBoards, convertBoard(board))
				board[i][j] = '.'
				break
			} else {
				solveNQueensHelper(board, i+1)
			}
			board[i][j] = '.'
		}
	}
}

func emptyBoard(n int) [][]byte {
	str := strings.Repeat(".",n)
	board := make([][]byte, n)
	for i := range board {
		board[i] = []byte(str)
	}
	return board
}

func convertBoard(board [][]byte) []string {
	result := make([]string,len(board))
	for i,row := range board {
		result[i] = string(row)
	}
	return result
}


type Point [2]int

func canPlaceQueen(board [][]byte, x,y int) bool {
	n := len(board)
	// Look for queens up and down
	for i := range x {
		if board[i][y] == 'Q' {
			return false
		}
	}

	// Look for queens on the angles
	point := Point{x-1,y-1}
	for {
		if point[0] < 0 || point[1] < 0 {
			break
		}
		if board[point[0]][point[1]] == 'Q' {
			return false
		}
		point = Point{point[0]-1, point[1]-1}
	}

	point = Point{x-1,y+1}
	for {
		if point[0] < 0 || point[1] >= n {
			break
		}
		if board[point[0]][point[1]] == 'Q' {
			return false
		}
		point = Point{point[0]-1, point[1]+1}
	}
	return true
}