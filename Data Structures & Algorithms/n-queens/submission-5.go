
var allBoards [][]string
func solveNQueens(n int) [][]string {
	allBoards = nil
	board := emptyBoard(n)
	solveNQueensHelper(board, 0)
	return allBoards
}

func solveNQueensHelper(board [][]byte, rowI int) {
	for j := range board {
		if canPlaceQueen(board, rowI, j) {
			board[rowI][j] = 'Q'
			if rowI == len(board)-1 {
				allBoards = append(allBoards, convertBoard(board))
			} else {
				solveNQueensHelper(board, rowI+1)
			}
			board[rowI][j] = '.'
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
	points := []Point{
		Point{x-1,y-1},
		Point{x-1,y+1},
	}
	// Look for queens on the angles
	for {
		var oneInBounds bool
		for _,point := range points {
			if point[0] < 0 || point[0] >=n {
				continue
			}
			if point[1] < 0 || point[1] >=n {
				continue
			}
			oneInBounds = true
			if board[point[0]][point[1]] == 'Q' {
				return false
			}
		}
		if !oneInBounds {
			break
		}
		points[0] = Point{ points[0][0]-1,points[0][1]-1}
		points[1] = Point{ points[1][0]-1,points[1][1]+1}
	}
	return true
}