
func checkInclusion(s1 string, s2 string) bool {

	counts1 := map[rune]int{}
	for _,c := range s1 {
		counts1[c]++
	}
	neededCounts := len(counts1)
	matchingCounts := 0
	// fmt.Printf("counts1 %v\n", counts1)
	counts2 := map[rune]int{}
	for right, c := range s2 {
		counts2[c]++
		if counts2[c]==counts1[c] {
			matchingCounts++
		}
		if right >= len(s1) {
			letterToRemove := rune(s2[right-len(s1)])
			counts2[letterToRemove]--
			if counts2[letterToRemove] == 0 {
				delete(counts2,letterToRemove)
			}
			if counts2[letterToRemove] == counts1[letterToRemove] -1 {
				matchingCounts--
			}
		}

		// fmt.Printf("right %v right-len(s1) %v counts2 %v\n", right,right-len(s1),counts2)
		if neededCounts == matchingCounts {
			return true
		}
	}
	return false
}
