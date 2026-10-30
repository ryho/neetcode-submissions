// 1:33
func dailyTemperatures(temps []int) []int {
	var stack []int
	output := make([]int, len(temps))
	for i, temp := range temps {
		if len(stack) > 0 && temp > temps[stack[len(stack)-1]] {
			for j := len(stack)-1;len(stack)>0;j-- {
				k := stack[j]
				if temps[k] >= temp {
					break
				}
				stack = stack[:len(stack)-1]
				output[k] = i-k
			}
		}
		stack = append(stack, i)
	}
	return output
}