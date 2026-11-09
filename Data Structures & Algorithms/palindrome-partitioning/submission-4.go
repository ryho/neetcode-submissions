import "slices"

/*
"aabaa"
[["a","a","b","a","a"],["aa","b","a","a"],["a","a","b","aa"],["aa","b","aa"],["aabaa"]]
*/
var allResults [][]string

func partition(s string) [][]string {
	allResults = nil
	palindromeMap = map[[2]int]bool{}
	partitionHelper(s, 0, nil)

	return allResults
}

func partitionHelper(s string, i int, current []string) {
	if i == len(s) {
		allResults = append(allResults, slices.Clone(current))
	}
	for j := len(s)-1; j>=i; j-- {
		if checkPalindrome(s, i, j) {
			partitionHelper(s, j+1, append(current, s[i:j+1]))
		}
	}
}
var palindromeMap map[[2]int]bool

func checkPalindrome(s string, i, j int) bool {
	p := [2]int{i,j}
	result, ok := palindromeMap[p]
	if ok {
		return result
	}
	for i<j {
		if s[i] != s[j] {
			palindromeMap[p] = false
			return false
		}
		i++
		j--
	}
	palindromeMap[p] = true
	return true
}