func isAnagram(s string, t string) bool {
	sData := createMap(s)
	tData := createMap(t)
	if len(sData) != len(tData) {
		return false
	}
	for k,v := range sData {
		if tData[k] != v {
			return false
		}
	}
	return true
}

func createMap(s string) map[rune]int {
	data := map[rune]int{}
	for _,s := range s {
		data[s]+=1
	}
	return data
}
