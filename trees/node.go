package trees

type Node struct {
	Value  int
	Left   *Node
	Right  *Node
	Parent *Node
}

func NewNode(value int) *Node {
	return &Node{value, nil, nil, nil}
}

func (node *Node) AddParent(parent *Node) {
	node.Parent = parent
}

func (node *Node) AddChildren(leftNode, rightNode *Node) {
	node.Left = leftNode
	node.Right = rightNode

	if leftNode != nil {
		leftNode.AddParent(node)
	}

	if rightNode != nil {
		rightNode.AddParent(node)
	}
}

func ExtractPathToRoot(node *Node) []int {
	path := []int{}

	if node == nil {
		return path
	}

	for {
		path = append([]int{node.Value}, path...)
		if node.Parent != nil {
			node = node.Parent
		} else {
			break
		}
	}

	return path
}
