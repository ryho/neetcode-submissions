
var allStrings []string

func generateParenthesis(n int) []string {
	allStrings = nil 
	generateParenthesisHelper(n, 0, "")
	return allStrings
}

func generateParenthesisHelper(n int, openCount int, current string) {
	if n == 0 && openCount == 0 {
		allStrings = append(allStrings, current)
		return
	}
	if n > 0 {
		currentOpen := current + "("
		generateParenthesisHelper(n-1,openCount+1,currentOpen)
	}
	if openCount > 0 {
		currentClosed := current + ")"
		generateParenthesisHelper(n,openCount-1, currentClosed)
	}
}
