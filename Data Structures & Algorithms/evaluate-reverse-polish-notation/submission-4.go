

func evalRPN(tokens []string) int {
	var stack []int
	for _, token := range tokens {
		if isOp(token) {
			if len(stack) < 2 {
				fmt.Printf("stack len %v\n", len(stack))
				return math.MinInt
			}
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			var res int
			if token == "+" {
				res = left+right
			} else if token == "-" {
				res = left-right
			} else if token == "*" {
				res = left*right
			} else if token == "/" {
				res = left/right
			} else {
				fmt.Printf("invalid op %v\n", token)
				return math.MinInt
			}
			stack = stack[:len(stack)-1]
			stack[len(stack)-1] = res
		} else {
			num, err := strconv.Atoi(token)
			if err != nil {
				fmt.Printf("Atoi err %v\n", err)
				return math.MaxInt
			}
			stack = append(stack, num)
		}
	}
	return stack[0]
}

func isOp(s string) bool{
	return len(s) == 1 &&(s == "-" || s == "+" || s == "*" || s == "/")
}