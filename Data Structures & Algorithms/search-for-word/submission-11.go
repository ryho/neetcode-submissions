
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
	path := map[Point]bool{}
	for i, row := range board {
		for j, _ := range row {
			p := Point{i, j}
			path[p] = true
			found := check(board, word, 0, p, path)
			if found {
				return true
			}
			delete(path, p)
		}
	}
	return false
}

func check(board [][]byte, word string, i int, p Point, path map[Point]bool) bool {
	// fmt.Printf("check word %v p %v, i %v\n", word, p, i)
	if index(board,p) == word[i] {
		if i == len(word)-1  {
			return true
		}
		path[p] = true
		i++
		a := Point{p[0]-1, p[1]}
		if validate(board, path, a) {
			if check(board, word, i, a, path) {
				return true
			}
		}
		b := Point{p[0]+1,p[1]}
		if validate(board, path, b) {
			if check(board, word, i, b, path) {
				return true
			}
		}
		c := Point{p[0],p[1]-1}
		if validate(board, path, c) {
			if check(board, word, i, c, path) {
				return true
			}
		}
		d := Point{p[0],p[1]+1}
		if validate(board, path, d) {
			if check(board, word, i, d, path) {
				return true
			}
		}
		delete(path, p)
	}
	return false
}

// index should only be called with points that exist in board
func index(board [][]byte, p Point) byte {
	return board[p[0]][p[1]]
}

func validate(board [][]byte, path map[Point]bool, p Point) bool {
	if !exists(board, p) {
		return false
	}
	if path[p] {
		return false
	}
	return true
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