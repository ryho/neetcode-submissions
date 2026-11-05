import "slices"
var matches [][]int
func combinationSum2(candidates []int, target int) [][]int {
	matches = nil
	slices.Sort(candidates)
	combinationSum2Helper(candidates, 0, target, 0, nil)
	return matches
}

func combinationSum2Helper(candidates []int, cIndex, target,currentCount int, currentList []int) {
	if cIndex >= len(candidates) {
		return
	}
	num := candidates[cIndex]

	if num > target-currentCount {
		return
	}
	// Try with out num
	combinationSum2Helper(candidates, cIndex+1, target, currentCount, currentList)
	

	// Trick: if the last number was the same number, and we skipped it,
	// do not add the number
	for temp := 1; cIndex - temp >= 0; temp++{
		if candidates[cIndex - temp] == num {
			if len(currentList) < temp || currentList[len(currentList)-temp] != num {
				return
			}
		}
	}

	// Try with num
	currentCount += num
	currentList = append(currentList, num)
	if currentCount == target {
		matches = append(matches, slices.Clone(currentList))
	} else if currentCount < target && len(candidates)-cIndex > 1 {
		combinationSum2Helper(candidates, cIndex+1, target, currentCount, currentList)
	}
}
