import(
	"slices"
)
func groupAnagrams(strs []string) [][]string {
	data := map[string]int{}
	var result [][]string
	for _, str := range strs {
		std := standardizeData(str)
		i, ok := data[std]
		if ok {
			result[i]=append(result[i], str)
		} else {
			data[std] = len(result)
			result = append(result, []string{str})
		}
	}
	return result
}

func standardizeData(str string) string{
	var data []string
	for i := range str {
		data = append(data, str[i:i+1])
	}
	slices.Sort(data)
	return strings.Join(data, "")
}