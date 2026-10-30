
func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	result[0] = nums[0]
	for i:=1;i<len(nums);i++ {
		result[i] = result[i-1]*nums[i]
	}

	cumulative := 1
	for i := len(nums)-1; i>0; i-- {
		result[i] = result[i-1]*cumulative
		cumulative *= nums[i]
	}
	result[0] = cumulative
	return result
}


func productExceptSelfDivision(nums []int) []int {
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
