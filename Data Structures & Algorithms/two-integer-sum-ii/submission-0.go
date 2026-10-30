
// 1 2 3 4 5 6 7 8 9 10
func twoSum(numbers []int, target int) []int {
	i := 0
	j := len(numbers)-1
	for {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i+1,j+1}
		} else if sum > target {
			j--
		} else if sum < target {
			i++
		}
		if i >= j {
			// did not find
			return nil
		}
	}
}
