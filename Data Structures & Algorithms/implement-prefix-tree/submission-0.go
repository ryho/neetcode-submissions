type PrefixTree struct {
	data map[byte]*PrefixTree
}

func Constructor() PrefixTree {
    return PrefixTree{
		data: map[byte]*PrefixTree{},
	}
}

func (this *PrefixTree) Insert(word string) {
	word = strings.ToLower(word)
	node := this
	for i := range word {
		c := word[i]
		next, ok := node.data[c]
		if !ok {
			next = &PrefixTree{
				data: map[byte]*PrefixTree{},
			}
			node.data[c] = next
		}
		node = next
	}
	node.data[0] = nil
}

func (this *PrefixTree) Search(word string) bool {
	word = strings.ToLower(word)
	node := this
	for i := range word {
		c := word[i]
		next, ok := node.data[c]
		if !ok {
			return false
		}
		node = next
	}
	_, ok := node.data[0]
	return ok
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	prefix = strings.ToLower(prefix)
	node := this
	for i := range prefix {
		c := prefix[i]
		next, ok := node.data[c]
		if !ok {
			return false
		}
		node = next
	}
	return true
}
