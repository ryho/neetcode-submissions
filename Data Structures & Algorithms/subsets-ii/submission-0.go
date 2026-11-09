import "slices"

var subsets [][]int
func subsetsWithDup(nums []int) [][]int {
	slices.Sort(nums)
	subsets = nil

	subsetsWithDupHelper(nums, nil)

	return subsets
}

// subsetsWithDupHelper must be called with sorted nums only
func subsetsWithDupHelper(nums []int, current []int) {
	if len(nums) == 0 {
		subsets = append(subsets, slices.Clone(current))
		return 
	}
	num := nums[0]
	var end int // end of duplicate numers
	for i := 1; i < len(nums); i++ {
		if nums[i] == num {
			end = i
		} else {
			break
		}
	}
	// if end+1 < len(nums) {
	// 	nums = nums[end+1:]
	// } else {
	// 	nums = nil
	// }
	subsetsWithDupHelper(nums[end+1:], current)
	for i := 0; i <= end; i++ {
		current = append(current, num)
		subsetsWithDupHelper(nums[end+1:], current)
	}
}
