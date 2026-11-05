type PrefixTree struct {
	// Key is either the next character in the word,
	// or 0 for end of string character.
	data [27]*PrefixTree
}

func Constructor() PrefixTree {
    return PrefixTree{
		data: [27]*PrefixTree{},
	}
}
const END_CHAR int = 'z'-'a'+1

func (this *PrefixTree) Insert(word string) {
	word = strings.ToLower(word)
	node := this
	for i := range word {
		c := word[i]
		next := node.data[c-'a']
		if next == nil {
			next = &PrefixTree{
				data: [27]*PrefixTree{},
			}
			node.data[c-'a'] = next
		}
		node = next
	}
	node.data[END_CHAR] = &PrefixTree{}
}

func (this *PrefixTree) Search(word string) bool {
	word = strings.ToLower(word)
	node := this
	for i := range word {
		c := word[i]
		next := node.data[c-'a']
		if next == nil {
			return false
		}
		node = next
	}
	end := node.data[END_CHAR]
	return end != nil
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	prefix = strings.ToLower(prefix)
	node := this
	for i := range prefix {
		c := prefix[i]
		next := node.data[c-'a']
		if next == nil {
			return false
		}
		node = next
	}
	return true
}
