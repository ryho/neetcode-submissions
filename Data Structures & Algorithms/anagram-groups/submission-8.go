import(
	"slices"
)
func groupAnagrams(strs []string) [][]string {
	data := map[string][]string{}
	for _, str := range strs {
		std := standardizeData(str)
		data[std] = append(data[std], str)
	}
	var result [][]string
	for _,strs := range data{
		result = append(result, strs)
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