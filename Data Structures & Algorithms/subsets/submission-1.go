func subsets(nums []int) [][]int {
	resultLen := int(math.Pow(2, float64(len(nums))))
	result := make([][]int, resultLen)

	for i, num := range nums {
		j := int(math.Pow(2,float64(i)))
		inARow := j
		jump := j
		inARowCount :=0
		for  j<resultLen {
			result[j] = append(result[j], num)
			inARowCount++
			if inARowCount == inARow {
				inARowCount = 0
				j += jump
			}
			j++
		}
	}

	return result
}
/*
0:1
1:2
2:4
3:8
*/