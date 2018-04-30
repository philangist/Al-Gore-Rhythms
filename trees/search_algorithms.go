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

	if rootNode.Left == nil && rootNode.Right == nil {
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

	if rootNode.Left == nil && rootNode.Right == nil {
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

		if currentNode.Left != nil {
			queue = append(queue, currentNode.Left)
		}

		if currentNode.Right != nil {
			queue = append(queue, currentNode.Right)
		}

		j += 1
	}

	return currentNode
}
