//3:56
//4:36
import (
"slices"
)
func threeSum(nums []int) [][]int {
	var result [][]int
	slices.Sort(nums)
	lastIVal := math.MinInt
	for i, num := range nums {
		// skip matching numbers
		if num == lastIVal {
			continue
		}
		lastIVal = num
		j:= i+1
		k := len(nums)-1
		for {
			// if the numbers meet, we are done looking
			if j >= k {
				break
			}
			sum := num + nums[j] + nums[k]
			if sum == 0 {
				result = append(result, []int{num, nums[j], nums[k]})
				kVal := nums[k]
				// skip matching numbers
				for {
					k--
					if nums[k]!=kVal || k == j || k==i {
						break
					}
				}
				jVal := nums[j]
				for {
					j++
					if nums[j]!=jVal || k == j || j==len(nums)-1{
						break
					}
				}
			} else if sum > 0 {
				k--
			} else if sum < 0 {
				j++
			}
		}
	}
	return result
}
