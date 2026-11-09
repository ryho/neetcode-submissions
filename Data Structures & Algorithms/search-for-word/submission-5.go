
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
			p := Point{i, j}
			found := check(board, word, p, map[Point]bool{p:true})
			if found {
				return true
			}
		}
	}
	return false
}

func check(board [][]byte, word string, p Point, path map[Point]bool) bool {
	// fmt.Printf("check word %v p %v\n", word, p)
	if index(board,p) == word[0] {
		if len(word) == 1 {
			return true
		}
		word = word[1:]
		a := Point{p[0]-1, p[1]}
		if validate(board, path, a) {
			path[a] = true
			if check(board, word, a, path) {
				return true
			}
			delete(path, a)
		}
		b := Point{p[0]+1,p[1]}
		if validate(board, path, b) {
			path[b] = true
			if check(board, word, b, path) {
				return true
			}
			delete(path, b)
		}
		c := Point{p[0],p[1]-1}
		if validate(board, path, c) {
			path[c] = true
			if check(board, word, c, path) {
				return true
			}
			delete(path, c)
		}
		d := Point{p[0],p[1]+1}
		if validate(board, path, d) {
			path[d] = true
			if check(board, word, d, path) {
				return true
			}
			delete(path, d)
		}
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