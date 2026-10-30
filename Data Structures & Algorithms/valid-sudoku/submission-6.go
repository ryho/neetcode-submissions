/* board=[
[".",".","4",".",".",".","6","3","."],
[".",".",".",".",".",".",".",".","."],
["5",".",".",".",".",".",".","9","."],
[".",".",".","5","6",".",".",".","."],
["4",".","3",".",".",".",".",".","1"],
[".",".",".","7",".",".",".",".","."],
[".",".",".","5",".",".",".",".","."],
[".",".",".",".",".",".",".",".","."],
[".",".",".",".",".",".",".",".","."]]
*/

func isValidSudoku(board [][]byte) bool {
	columnMaps := make([]map[byte]bool, 9)
	boxMaps := make([]map[byte]bool, 9)
	for i := range columnMaps {
		columnMaps[i] = map[byte]bool{}
		boxMaps[i] = map[byte]bool{}
	}
	for i,row := range board {
		rowMap := map[byte]bool{}
		for j, cell := range row {
			if cell != '.' {
				_, ok := rowMap[cell]
				if ok {
					return false
				}
				rowMap[cell] = true

				_, ok = columnMaps[j][cell]
				if ok {
					return false
				}
				columnMaps[j][cell] = true

				boxIndex := calcBoxIndex(i,j)
				_, ok = boxMaps[boxIndex][cell]
				if ok {
					return false
				}
				boxMaps[boxIndex][cell] = true
			}
		}
	}
	return true
}

func calcBoxIndex(i,j int) int{
	return boxIndexHelper(i)+boxIndexHelper(j)*3
}

func boxIndexHelper(i int) int{
	if i<3 {
		return 0
	} else if i>=3 && i< 6 {
		return 1
	}
	return 2
}
