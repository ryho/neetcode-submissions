import "maps"
func checkInclusion(s1 string, s2 string) bool {

	counts1 := map[rune]int{}
	for _,c := range s1 {
		counts1[c]++
	}
	// fmt.Printf("counts1 %v\n", counts1)
	counts2 := map[rune]int{}
	for right, c := range s2 {
		counts2[c]++
		if right >= len(s1) {
			letterToRemove := rune(s2[right-len(s1)])
			counts2[letterToRemove]--
			if counts2[letterToRemove]==0 {
				delete(counts2,letterToRemove)
			}
		}

		// fmt.Printf("right %v right-len(s1) %v counts2 %v\n", right,right-len(s1),counts2)
		if maps.Equal(counts1, counts2) {
			return true
		}
	}
	return false
}
