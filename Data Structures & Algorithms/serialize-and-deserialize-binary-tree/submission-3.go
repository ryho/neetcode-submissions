/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    
}

func Constructor() Codec {
    return Codec{}
}
const rowConst string = "row"

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
    queue := []*TreeNode{root}
	rowLength := len(queue)
	var currentRow []int
	allData :=[][]int{[]int{root.Val}}
	var i int
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil {
			currentRow = append(currentRow, node.Left.Val)
			queue = append(queue, node.Left)
		} else {
			currentRow = append(currentRow, math.MinInt)
		}
		if node.Right != nil {
			currentRow = append(currentRow, node.Right.Val)
			queue = append(queue, node.Right)
		} else {
			currentRow = append(currentRow, math.MinInt)
		}
		i++
		if i == rowLength {
			i = 0
			allData = append(allData, currentRow)
			currentRow = nil
			rowLength = len(queue)
		}
	}
	return serializeArrayOfArrays(allData)
}

func serializeArrayOfArrays(input [][]int) string{
	var rows []string
	for _, row := range input {
		var rowStrings []string
		allNull := true
		for _, cell := range row {
			if cell == math.MinInt {
				rowStrings = append(rowStrings, "null")
			} else {
				allNull = false
				rowStrings = append(rowStrings, strconv.Itoa(cell))
			}
		}
		if !allNull{
			rows = append(rows, strings.Join(rowStrings, ","))
		}
	}
	return strings.Join(rows, rowConst)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(dataStr string) *TreeNode {
	// fmt.Printf("data %v\n", dataStr)
	if dataStr == "" {
		return nil
	}
	rows := strings.Split(dataStr, rowConst)
	var data [][]int
	for _, row := range rows {
		cells := strings.Split(row, ",")
		var row []int
		for _, c := range cells {
			if c == "null" {
				row = append(row,math.MinInt)
			} else {
				num, err := strconv.Atoi(c)
				if err != nil {
					panic(err)
				}
				row = append(row,num)
			}
		}
		data = append(data, row)
	}

	head := &TreeNode{
		Val: data[0][0],
	}
	i := 1
	j := 0
	queue := []*TreeNode{head}
	for len(queue) > 0 {
		if i == len(data) {
			break
		}
		node := queue[0]
		queue = queue[1:]
		// fmt.Printf("i %v, j %v, data %v\n", i ,j, data)
		left := data[i][j]
		j++
		right := data[i][j]
		j++
		if j == len(data[i]) {
			i++
			j=0
		}
		if left != math.MinInt {
			leftNode := &TreeNode{
				Val: left,
			}
			node.Left = leftNode
			queue = append(queue, leftNode)
		}
		if right != math.MinInt {
			rightNode := &TreeNode{
				Val: right,
			}
			node.Right = rightNode
			queue = append(queue, rightNode)
		}
	}

	return head
}

