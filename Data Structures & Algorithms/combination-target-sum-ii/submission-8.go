import "slices"
var matches [][]int
func combinationSum2(candidates []int, target int) [][]int {
	matches = nil
	slices.Sort(candidates)
	fmt.Printf("candidates %v\n",candidates)
	combinationSum2Helper(candidates, 0, target, 0, nil)
	return matches
}

func combinationSum2Helper(candidates []int, cIndex, target,currentCount int, currentList []int) {
	if cIndex >= len(candidates) {
		return
	}
	// Try with out num
	combinationSum2Helper(candidates,cIndex+1, target, currentCount, currentList)
	
	// Try with num
	num := candidates[cIndex]

	// Trick: if the last number was the same number, and we skipped it,
	// do not add the number
	for temp := 1; cIndex - temp >= 0; temp++{
		if candidates[cIndex - temp] == num {
			if len(currentList) >= temp {
				if currentList[len(currentList)-temp] != num {
					return
				}
			} else {
				return
			}
		}
		
	}
	if cIndex > 0 {
		if candidates[cIndex-1] == num {
			if len(currentList) > 0 {
				if currentList[len(currentList)-1] != num {

				}
			}
		}
	}


	currentCount += num
	currentList = append(currentList, num)
	if currentCount == target {
		matches = append(matches, slices.Clone(currentList))
	} else if currentCount < target && len(candidates)-cIndex > 1 {
		combinationSum2Helper(candidates, cIndex+1, target, currentCount, currentList)
	}
}
