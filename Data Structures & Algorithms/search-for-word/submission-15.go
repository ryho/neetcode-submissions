
type Point [2]int

/*
board=[
["A","B","C","E"],
["S","F","E","S"],
["A","D","E","E"]]
word="ABCESEEEFS"
*/
func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	path := map[Point]bool{}
	for i, row := range board {
		for j, _ := range row {
			p := Point{i, j}
			found := searchForWord(board, word, 0, p, path)
			if found {
				return true
			}
		}
	}
	return false
}

func searchForWord(board [][]byte, word string, i int, p Point, path map[Point]bool) bool {
	// fmt.Printf("check word %v p %v, i %v\n", word, p, i)
	if index(board,p) == word[i] {
		// If we've matche the whole word now, return true.
		if i == len(word)-1  {
			return true
		}
		path[p] = true
		i++
		points := []Point{
			Point{p[0]-1, p[1]},
			Point{p[0]+1, p[1]},
			Point{p[0], p[1]-1},
			Point{p[0], p[1]+1},
		}
		for _, point := range points {
			if validate(board, path, point) && searchForWord(board, word, i, point, path) {
				return true
			}
		}
		// Remove points from the path when backtracking
		delete(path, p)
	}
	return false
}

// index should only be called with points that exist in board
func index(board [][]byte, p Point) byte {
	return board[p[0]][p[1]]
}

// validate checks that the point is within the bounds of the board
// and that is not already on the current path
func validate(board [][]byte, path map[Point]bool, p Point) bool {
	if !exists(board, p) {
		return false
	}
	if path[p] {
		return false
	}
	return true
}

// Checks if the indexes exists in the map
func exists(board [][]byte, p Point) bool {
	if p[0] < 0 || p[0] >= len(board) {
		return false
	}
	if p[1] < 0 || p[1] >= len(board[0]) {
		return false
	}
	return true
}