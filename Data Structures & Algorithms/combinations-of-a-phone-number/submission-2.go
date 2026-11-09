
var allStrings []string = make([]string, 0, int(math.Pow(4,4)))
func letterCombinations(digits string) []string {
	allStrings = allStrings[:0]
	if len(digits) == 0 {
		return allStrings
	}
	letterCombinationsHelper(digits, 0, make([]byte,0,len(digits)))
	return allStrings
}

func letterCombinationsHelper(digits string, i int, letters []byte) {
	if i == len(digits) {
		allStrings = append(allStrings, string(letters))
		return
	}
	possible := numsToLetters[digits[i]-'2']
	i++
	for _, char := range possible {
		letterCombinationsHelper(digits, i, append(letters, char))
	}
}


var numsToLetters = [][]byte {
	[]byte{'a','b','c'},
	[]byte{'d','e','f'},
	[]byte{'g','h','i'},
	[]byte{'j','k','l'},
	[]byte{'m','n','o'},
	[]byte{'p','q','r','s'},
	[]byte{'t','u','v'},
	[]byte{'w','x','y','z'},
}