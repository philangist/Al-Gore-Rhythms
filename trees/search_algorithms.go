package trees

import "fmt"

func DepthFirstSearch(rootNode *Node, value int) *Node {
    if rootNode == nil {
        return nil
    }

    fmt.Println("Visiting node ", rootNode.Value)
    if rootNode.Value == value {
        fmt.Println("Found value ", rootNode.Value)
        return rootNode
    }

    if (rootNode.Left == nil && rootNode.Right == nil){
        return nil
    }

    leftBranch := DepthFirstSearch(rootNode.Left, value)
    if leftBranch != nil {
        return leftBranch
    }
    return DepthFirstSearch(rootNode.Right, value)
}

func BreadthFirstSearch(rootNode *Node, value int) *Node {
    if rootNode == nil {
        return nil
    }

    if rootNode.Value == value {
        fmt.Println("Found value ", rootNode.Value)
        return rootNode
    }

    if (rootNode.Left == nil && rootNode.Right == nil){
        return nil
    }

    queue := []*Node{rootNode}
    var currentNode *Node

    j := 0
    for {
        if j == len(queue) {
            return nil
        }
        currentNode = queue[j]
        fmt.Println("Visiting node ", currentNode.Value)

        if currentNode.Value == value {
            fmt.Println("Found value ", currentNode.Value)
            return currentNode
        }

        if (currentNode.Left != nil){
            queue = append(queue, currentNode.Left)
        }

        if (currentNode.Right != nil){
            queue = append(queue, currentNode.Right)
        }

        j += 1
    }

    return currentNode
}
/*
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
*/