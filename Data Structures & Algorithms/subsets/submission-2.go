func subsets(nums []int) [][]int {
	resultLen := int(math.Pow(2, float64(len(nums))))
	result := make([][]int, resultLen)

	for i, num := range nums {
		sequenceLength := int(math.Pow(2,float64(i)))
		inARowCount := 0
		for j:= sequenceLength; j<resultLen; j++ {
			result[j] = append(result[j], num)
			inARowCount++
			if inARowCount == sequenceLength {
				inARowCount = 0
				j += sequenceLength
			}
		}
	}
	return result
}