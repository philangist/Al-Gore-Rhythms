package trees

import (
    "sort"
)

func NewRootNode(value int) *Node {
    return &Node{value, nil, nil, nil}
}

func NewSearchTree(values []int) *Node {
    if values == nil {
        return nil
    }

    sort.Slice(
        values, func(i, j int) bool {
            return values[i] < values[j]
        },
    )

    return ConstructSearchTree(values)
}

func ConstructSearchTree(values []int) *Node {
    if values == nil {
        return nil
    }

    if values == nil {
        return nil
    }

    lengthValues := len(values)

    if lengthValues == 1 {
        return NewRootNode(values[0])
    }

    pivotIndex := lengthValues / 2
    pivotValue := values[pivotIndex]
    pivotNode := NewRootNode(pivotValue)

    var leftNode, rightNode *Node

    leftValues := make([]int, pivotIndex)
    copy(leftValues, values[:pivotIndex])
    leftNode = ConstructSearchTree(leftValues)

    if pivotIndex < lengthValues {
        rightValues := make([]int, lengthValues - pivotIndex - 1)
        copy(rightValues, values[pivotIndex + 1:])
        rightNode = ConstructSearchTree(rightValues)
    }

    pivotNode.AddChildren(leftNode,rightNode)

    return pivotNode
}

/*func (n *Node) Insert(value int) error { return nil }

func (n *Node) Search(value int) error { return nil }

func (n *Node) Delete(value int) error { return nil } */