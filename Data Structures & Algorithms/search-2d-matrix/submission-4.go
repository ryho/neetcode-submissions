func searchMatrix(matrix [][]int, target int) bool {
	left := 0
	right := (len(matrix)*len(matrix[0]))-1
	for {
		if right <= left {
			return indexMatrix(matrix, left) == target
		}
		index := left + (right-left)/2
		value := indexMatrix(matrix, index)
		if value == target {
			return true
		}else if value > target{
			right = index-1
		} else {
			left = index + 1
		}
	}
	return true
}

func indexMatrix(matrix [][]int, i int) int {
	x:= i%len(matrix[0])
	y:= i/len(matrix[0])
	return matrix[y][x]
}
