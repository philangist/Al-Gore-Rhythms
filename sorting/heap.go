package sorting

import (
    "fmt"
    "math"
)


const NIL_VALUE = -1
var NIL_NODE = NewHeapNode(NIL_VALUE, NIL_VALUE)

type Heap struct{
    Values []int
}

func NewHeap(values []int) *Heap {
    return &Heap{Values: values}
}

type HeapNode struct {
    Index int
    Value int
}

func NewHeapNode(index, value int) *HeapNode {
    return &HeapNode{Index: index, Value: value}
}

func (h *Heap) Length() int {
    return len(h.Values)
}

func (h *Heap) Parent(node *HeapNode) *HeapNode {
    return h.GetNode(int(math.Floor(float64(node.Index) / 2.0)))
}

func (h *Heap) GetNode(index int) *HeapNode {
    if index == NIL_VALUE {
        return NIL_NODE
    }
    if index > h.Length() - 1{
        return NIL_NODE
    }
    return NewHeapNode(index, h.Values[index])
}

func (h *Heap) GetRootNode() *HeapNode {
    heapRoot := h.GetNode(0)
    return heapRoot
}

func (h *Heap) GetChildren(node *HeapNode) []*HeapNode {
    if node.Index == NIL_VALUE {
        return []*HeapNode{}
    }
    leftChildIndex, rightChildIndex := (2 * (node.Index) + 1), (2 * (node.Index) + 2)
    if leftChildIndex > h.Length() {
        return []*HeapNode{NIL_NODE, NIL_NODE}
    }
    leftChild := h.GetNode(leftChildIndex)
    if rightChildIndex > h.Length() {
        return []*HeapNode{leftChild, NIL_NODE}
    }
    rightChild := h.GetNode(rightChildIndex)
    return []*HeapNode{leftChild, rightChild}
}

func (h *Heap)assertInvariant() error {
    index := 0
    nodes := []*HeapNode{h.GetRootNode()}

    for {
        node := nodes[index]
        children := h.GetChildren(node)
        leftChild, rightChild := children[0], children[1]
        if (leftChild.Value != NIL_VALUE){
            if (leftChild.Value > node.Value){
                return fmt.Errorf(
                    "Heap invariant violated for node at %d by child at %d",
                    node.Index, leftChild.Index,
                )
            }
            nodes = append(nodes, leftChild)
        }
        if (rightChild.Value != NIL_VALUE){
            if (rightChild.Value > node.Value){
                return fmt.Errorf(
                    "Heap invariant violated for node at %d by child at %d",
                    node.Index, rightChild.Index,
                )
            }
            nodes = append(nodes, rightChild)
        }

        index += 1
        if index > len(nodes) - 1 {
            break
        }
    }
    return nil
}

func ReorderHeap(h * Heap){}

func (h *Heap) PopHeapRoot() *HeapNode {
    heapRoot := h.GetRootNode()
    h.Values = h.Values[1:]
    ReorderHeap(h)
    return heapRoot
}