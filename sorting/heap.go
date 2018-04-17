package sorting

const NIL = -1

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
    return h.GetNode(node.Index / 2)
}

func (h *Heap) GetNode(index int) *HeapNode {
    if index == NIL {
        return nil
    }
    return NewHeapNode(index, h.Values[index])
}

func (h *Heap) GetChildren(node *HeapNode) []*HeapNode {
    if node.Index == NIL {
        return nil
    }
    leftChildIndex, rightChildIndex := (2 * (node.Index) + 1), (2 * (node.Index) + 2)
    if leftChildIndex > h.Length() {
        return []*HeapNode{nil, nil}
    }
    leftChild := h.GetNode(leftChildIndex)
    if rightChildIndex > h.Length() {
        return []*HeapNode{leftChild, nil}
    }
    rightChild := h.GetNode(rightChildIndex)
    return []*HeapNode{leftChild, rightChild}
}