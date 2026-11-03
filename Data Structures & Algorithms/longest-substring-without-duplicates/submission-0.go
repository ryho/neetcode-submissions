func lengthOfLongestSubstring(s string) int {
	check := map[rune]int{}
	start := 0
	maxSize := 0
	for i, c:= range s {
		index, ok := check[c]
		if ok {
			for j := start;j<=index;j++ {
				delete(check, rune(s[j]))
			}
			start = index+1
		}
		check[c] = i
		if len(check) > maxSize {
			maxSize = len(check)
		}
	}
	return maxSize
}
