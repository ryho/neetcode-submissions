import "maps"

type Point [2]int

/*
board=[
["A","B","C","E"],
["S","F","E","S"],
["A","D","E","E"]]
word="ABCESEEEFS"
*/
func exist(board [][]byte, word string) bool {
	if word == "" {
		return true
	}
	for i, row := range board {
		for j, _ := range row {
			found := check(board, word, Point{i, j}, map[Point]bool{})
			if found {
				return true
			}
		}
	}
	return false
}

func check(board [][]byte, word string, p Point, path map[Point]bool) bool {
	// fmt.Printf("check word %v p %v\n", word, p)
	if !exists(board, p) {
		return false
	}
	if path[p] {
		return false
	}
	if index(board,p) == word[0] {
		if len(word) == 1 {
			return true
		}
		path[p] = true
		word = word[1:]
		a := Point{p[0]-1, p[1]}
		if check(board, word, a, maps.Clone(path)) {
			return true
		}
		b := Point{p[0]+1,p[1]}
		if check(board, word, b, maps.Clone(path)) {
			return true
		}
		c := Point{p[0],p[1]-1}
		if check(board, word, c, maps.Clone(path)) {
			return true
		}
		d := Point{p[0],p[1]+1}
		if check(board, word, d, maps.Clone(path)) {
			return true
		}
	}
	return false
}

// index should only be called with points that exist in board
func index(board [][]byte, p Point) byte {
	return board[p[0]][p[1]]
}

func exists(board [][]byte, p Point) bool {
	if p[0] < 0 || p[0] >= len(board) {
		return false
	}
	if p[1] < 0 || p[1] >= len(board[0]) {
		return false
	}
	return true
}