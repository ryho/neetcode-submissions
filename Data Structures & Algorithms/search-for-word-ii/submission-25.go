// import "slices"

var foundWords []string = []string{}

var boardToLog [][]byte = [][]byte{[]byte{'a'}}
var loggingOn bool
func findWords(board [][]byte, words []string) []string {
	tree := NewTree(words...)

	// if slices.EqualFunc(board, boardToLog, slices.Equal) {
	// 	loggingOn = true
	// }

	if loggingOn {
		fmt.Printf("tree.data %v\n", tree.data)
	}

	foundWords = nil
	for i, row := range board {
		for j, cell := range row {
			nextTree, ok := tree.FindPrefixByte(cell) 
			if ok {
				firstPoint := [2]int{i,j}
				indexes := map[[2]int]bool{
					firstPoint: true,
				}
				word := []byte{cell}
				_, ok := nextTree.FindPrefixByte(0)
				if ok {
					foundWords = append(foundWords, string(word))
					nextTree.Delete(0)
				}
				findNextCharacter(board,nextTree,indexes,firstPoint,word)
			}
		}
	}

    return foundWords
}

func findNextCharacter(board [][]byte, tree *Tree, indexes map[[2]int]bool, last [2]int, word []byte) {
	if loggingOn {
		fmt.Printf("findNextCharacter word %v\n",word)
	}
	i, j := last[0], last[1]
	a := [2]int{i-1,j}
	b := [2]int{i+1,j}
	c := [2]int{i,j-1}
	d := [2]int{i,j+1}
	checkPoint(board,tree,indexes,last,a,word)
	checkPoint(board,tree,indexes,last,b,word)
	checkPoint(board,tree,indexes,last,c,word)
	checkPoint(board,tree,indexes,last,d,word)
}
func checkPoint(board [][]byte, tree *Tree, indexes map[[2]int]bool, last, next [2]int, word []byte) {
	if isPointValid(next, board, indexes) {
		nextTree, ok := tree.FindPrefixByte(board[next[0]][next[1]])
		if ok {
			word = append(word, board[next[0]][next[1]])
			_, ok := nextTree.FindPrefixByte(0)
			if ok {
				foundWords = append(foundWords, string(word))
				nextTree.Delete(0)
			}

			indexes[next]=true
			findNextCharacter(board, nextTree, indexes, next, word)
			indexes[next]=false
		}
	}
}

func isPointValid(point [2]int, board[][]byte, indexes map[[2]int]bool) bool {
	if indexes[point] {
		return false 
	}
	if point[0] < 0 || point[0] >= len(board) {
		return false
	}
	if point[1] < 0 || point[1] >= len(board[0]) {
		return false
	}
	return true
}

type Tree struct {
	data [27]*Tree
}

const END_CHAR = 'z'-'a'+1

func NewTree(words ...string) *Tree {
	t := &Tree{
		data:[27]*Tree{},
	}
	for _, word := range words {
		t.AddWord(word)
	}
	return t
}

func (t *Tree) AddWord(word string) {
	tree := t
	for i := range word {
		c := word[i]
		nextTree := tree.data[c-'a']
		if nextTree == nil {
			nextTree = NewTree()
			tree.data[c-'a'] = nextTree
		}
		tree = nextTree
	}
	tree.data[END_CHAR] = &Tree{}
}

func (t *Tree) FindWord(word string) bool {
	if len(word) == 0 {
		return t.data[END_CHAR] != nil
	}

	node := t.data[word[0]-'a']
	if node == nil {
		return false
	}
	return node.FindWord(word[1:])
}

func (t *Tree) FindPrefix(word string) (*Tree, bool) {
	if len(word) == 0 {
		return t, true
	}

	node := t.data[word[0]-'a']
	if node == nil {
		return nil, false
	}
	return node.FindPrefix(word[1:])
}

func (t *Tree) FindPrefixByte(c byte) (*Tree, bool) {
	if c == 0 {
		c = END_CHAR
	} else {
		c-='a'
	}
	tree := t.data[c]
	return tree, tree!=nil
}

func (t *Tree) Delete(c byte) {
	if c == 0 {
		c = END_CHAR
	} else {
		c-='a'
	}
	t.data[c] = nil
}
