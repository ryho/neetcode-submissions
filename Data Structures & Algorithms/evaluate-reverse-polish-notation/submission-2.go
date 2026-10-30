

func evalRPN(tokens []string) int {
	var stack []int
	for _, token := range tokens {
		if isNumeric(token) {
			num, err := strconv.Atoi(token)
			if err != nil {
				fmt.Printf("Atoi err %v\n", err)
				return math.MaxInt
			}
			stack = append(stack, num)
		} else {
			if len(stack) < 2 {
				fmt.Printf("stack len %v\n", len(stack))
				return math.MinInt
			}
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack)-1]
			if token == "+" {
				stack[len(stack)-1] = left+right
			} else if token == "-" {
				stack[len(stack)-1] = left-right
			} else if token == "*" {
				stack[len(stack)-1] = left*right
			} else if token == "/" {
				stack[len(stack)-1] = left/right
			} else {
				fmt.Printf("invalid op %v\n", token)
				return math.MinInt
			}
		}
	}
	return stack[0]
}

func isNumeric(s string) bool{
	if s == "-" {
		return false
	}
	for _, c := range s {
		if !(c >='0' && c <= '9' || c == '-') {
			return false
		}
	}
	return true
}