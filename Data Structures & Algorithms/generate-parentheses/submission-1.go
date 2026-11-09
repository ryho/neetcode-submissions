
var allStrings []string

func generateParenthesis(n int) []string {
	allStrings = nil 
	generateParenthesisHelper(n, 0, &strings.Builder{})
	return allStrings
}

func generateParenthesisHelper(n int, openCount int, current *strings.Builder) {
	if n == 0 && openCount == 0 {
		allStrings = append(allStrings, current.String())
		return
	}
	var cachedString string
	if n > 0 {
		if openCount > 0 {
			cachedString = current.String()
		}
		current.WriteByte('(')
		generateParenthesisHelper(n-1,openCount+1, current)
	}
	if openCount > 0 {
		if n > 0 { 
			sb := &strings.Builder{}
			sb.WriteString(cachedString)
			sb.WriteByte(')')
			generateParenthesisHelper(n,openCount-1, sb)
		} else {
			current.WriteByte(')')
			generateParenthesisHelper(n,openCount-1, current)
		}
	}
}
