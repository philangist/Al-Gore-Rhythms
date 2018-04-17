package sorting

// O(nlogn) read/write!
func HeapSort(input []int) []int {
    // heap sort
    heap := NewHeap(input)
    heap.ReorderHeap()

    output := []int{}
    for {
        heapRoot, _ := heap.PopRoot()
        if heapRoot == NIL_NODE {
            break
        }
        output = append([]int{heapRoot.Value}, output...)
    }
    return output
}