func isValid(s string) bool {
    var stack []rune
	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, c)
		} else{
			if len(stack) == 0 {
				return false
			}
			var expChar rune
			if c == ')'{
				expChar = '('
			} else if c == '}' {
				expChar = '{'
			} else if c == ']' {
				expChar = '['
			} else {
				return false
			}
			if stack[len(stack)-1] != expChar {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
