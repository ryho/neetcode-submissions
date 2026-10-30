func searchMatrix(matrix [][]int, target int) bool {
	left := 0
	right := (len(matrix)*len(matrix[0]))-1
	for {
		// fmt.Printf("left %v, right %v, target %v\n", left, right, target)
		if right <= left {
			if right >=0 {
				if indexMatrix(matrix, right) == target {
					return true
				}
			}
			return indexMatrix(matrix, left) == target
		}
		index := left + (right-left)/2
		value := indexMatrix(matrix, index)
		// fmt.Printf("index %v, value %v\n", index, value)
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
