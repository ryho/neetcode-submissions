func productExceptSelf(nums []int) []int {
	allProduct := 1
	var zeroCount int
	for _, num := range nums {
		if num == 0 {
			zeroCount++
			continue
		}
		allProduct *= num
	}
	if zeroCount > 1 {
		return make([]int, len(nums))
	}
	result:= make([]int, 0, len(nums))
	for _, num := range nums {
		if num == 0 {
			result = append(result, allProduct)
		} else if zeroCount == 1 {
			result = append(result, 0)
		} else {
			result = append(result, allProduct/num)
		}
	}
	return result
}
