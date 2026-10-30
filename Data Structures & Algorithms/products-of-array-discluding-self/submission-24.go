
func productExceptSelf(nums []int) []int {
	// nums 1 2 3 4  5   6   7    8     9 
	// res  1 1 2 6 24 120 720 5040 40320
	// res                     5040*9
	// res                 720*9*8


	// nums 1 2 4 6
	// res  1 2 8 8
	// res  1 2 12 8
	// 6

	// nums [-1,0,1,2,3]
	// res [1 ,]

	result := make([]int, len(nums))
	for i, num := range nums {
		if i == 0 {
			result[i] = num
		} else {
			result[i]= result[i-1]*num
		}
	}

	cumulative := 1
	for i := len(nums)-1; i>=0; i-- {
		if i == 0 {
			result[i] = cumulative
		} else {
			result[i] = result[i-1]*cumulative
		}
		cumulative *= nums[i]
	}
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
