func twoSum(nums []int, target int) []int {
	data:= map[int]int{}
	for i,v := range nums {
		j, ok := data[v]
		if ok {
			if i > j{
				return []int{j,i}
			}
			return []int{i,j}
		}
		data[target-v] = i
	}
	return []int{}
}

func twoSumOriginal(nums []int, target int) []int {
	var i,j int
    for x, xValue:= range nums {
		for y,yValue := range nums{
			if x==y{
				continue
			}
			if xValue + yValue == target {
				i=x
				j=y 
			}
		}
	}
	if i>j{
		return []int{j,i}
	}
	return []int{i,j}
}
