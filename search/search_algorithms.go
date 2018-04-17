package search

import "fmt"

type Node struct {
    value int
    left *Node
    right *Node
}

func NewNode (value int) *Node {
    return &Node{value, nil, nil}
}

func (node *Node) AddChildren (leftNode, rightNode *Node) {
    node.left = leftNode
    node.right = rightNode
}

func DepthFirstSearch(rootNode *Node, value int) *Node {
    fmt.Println("Visiting node ", rootNode.value)
    if rootNode.value == value {
        fmt.Println("Found value ", rootNode.value)
        return rootNode
    }

    if (rootNode.left == nil && rootNode.right == nil){
        return nil
    }

    leftBranch := DepthFirstSearch(rootNode.left, value)
    if leftBranch != nil {
        return leftBranch
    }
    return DepthFirstSearch(rootNode.right, value)
}

func main() {
    node1 := NewNode(1)
    node2 := NewNode(2)
    node3 := NewNode(3)
    node4 := NewNode(4)
    node5 := NewNode(5)
    node6 := NewNode(6)
    node7 := NewNode(7)
    node8 := NewNode(8)

    // use AddChildren

    node1.AddChildren(node2, node7)
    node2.AddChildren(node3, node4)
    node4.AddChildren(node5, node6)
    node7.AddChildren(node8, nil)

    DepthFirstSearch(node1, 4)
}
