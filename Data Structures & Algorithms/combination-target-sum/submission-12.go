import "slices"

var allLists [][]int
func combinationSum(nums []int, target int) [][]int {
	allLists = nil
	combinationSumHelper(nums, 0, target, nil)
	return allLists
}

func combinationSumHelper(nums []int, currentTotal, target int, currentList []int) {
	num := nums[0]
	runs := (target-currentTotal)/num
	for i := 0; i <= runs; i++ {
		if i != 0 {
			currentTotal += num
			currentList = append(currentList, num)
			if currentTotal == target {
			// fmt.Printf("found %v \n",currentList)
				allLists = append(allLists, slices.Clone(currentList))
			}
		}
		// fmt.Printf("combinationSumHelper num %v i %v runs %v currentTotal %v target %v list %v \n", num, i , runs, currentTotal, target, currentList)
		if currentTotal < target && len(nums) > 1 {
			combinationSumHelper(nums[1:], currentTotal,target,currentList)
		}
	}
}
