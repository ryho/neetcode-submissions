// 1:33
/*
Input: temperatures = [30,38,30,36,35,40,28]
out  1 0 0
Output: [1,4,1,2,1,0,0]

in 1 2 3 4 5 6 7 8 9 10
in 10 9 8 7 6 5 4 3 2 1 
*/
func dailyTemperatures(temps []int) []int {
	output := make([]int, len(temps))
	if len(temps) < 2 {
		return output
	}
	for i:=len(temps)-2;i>=0;i-- {
		currentTemp := temps[i]
		neighborIndex := i+1
		for {
			nextTemp := temps[neighborIndex]
			if nextTemp > currentTemp {
				output[i] = neighborIndex - i
				break
			} else {
				jump := output[neighborIndex]
				if jump == 0 {
					break
				}
				neighborIndex += jump
			}
		}
	}
	return output
}

func dailyTemperaturesStack(temps []int) []int {
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