//3:56
import (
"slices"
// "strconv"
)
func threeSum(nums []int) [][]int {
	var result [][]int
	slices.Sort(nums)
	// fmt.Printf("%v\n",nums)
	for i, num := range nums {
		j:= i+1
		k := len(nums)-1
		for {
			// if the numbers meet, we are done looking
			if j >= k {
				// fmt.Printf("jk meet %v %v %v\n",i,j,k)
				break
			}
			sum := num + nums[j] + nums[k]
			// fmt.Printf("sum %v\n",sum)
			if sum == 0 {
				result = append(result, []int{num, nums[j], nums[k]})
				k--
				j++
			} else if sum > 0 {
				k--
			} else if sum < 0 {
				j++
			}
		}
	}
	unique := map[string]bool{}
	var resultFiltered [][]int
	for _, trio := range result {
		str := strconv.Itoa(trio[0])+","+strconv.Itoa(trio[1])+","+strconv.Itoa(trio[2])
		if _, ok := unique[str]; ok {
			continue
		}
		resultFiltered = append(resultFiltered, trio)
		unique[str] = true
	}
	return resultFiltered
}
