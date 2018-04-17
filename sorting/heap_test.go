package sorting

import (
    "fmt"
    "testing"
)

func TestHeapCreation(t *testing.T){
    fmt.Println("Running TestHeapCreation...")
    heapSlice := []int{1, 2, 3, 4, 5}
    newHeap := NewHeap(heapSlice)
    if newHeap.Length() != len(heapSlice) {
        t.Errorf(
            "Mismatched length between created Heap %v and backing Slice %v",
            newHeap, heapSlice,
        )
    }
    for index, value := range heapSlice {
        heapNode := newHeap.GetNode(index)
        if index != heapNode.Index {
            t.Errorf(
                "HeapNode %v has index: %d.\nExpected index was: %d",
                heapNode, heapNode.Index, index,
            )
        }
        if value != heapNode.Value {
            t.Errorf(
                "HeapNode %v has value: %d.\nExpected value was: %d",
                heapNode, heapNode.Value, value,
            )
        }
    }
}



type heapChildrenTestCases struct {
    heapSlice []int
    parentsToChildren map[int][]int //parent index -> (leftChildIndex, rightChildIndex)
}

func compareChildren(expectedChildren []int, actualChildren []*HeapNode) bool {
    fmt.Println("actualChildren[0]: ", actualChildren[0])
    fmt.Println("actualChildren[1]: ", actualChildren[1])

    return ((expectedChildren[0] == actualChildren[0].Value) &&
            (expectedChildren[1] == actualChildren[1].Value))
}

func TestHeapChildren(t *testing.T){
    fmt.Println("Running TestHeapChildren...")
    testCases := []heapChildrenTestCases{
        {
            heapSlice: []int{4, 9, 7},
            parentsToChildren: map[int][]int{
                4: []int{9, 7},
            },
        },
        {
            heapSlice: []int{16, -1, -1},
            parentsToChildren: map[int][]int{
                16: []int{NIL, NIL},
            },
        },
    }

    for _, testCase := range testCases {
        heap := NewHeap(testCase.heapSlice)

        for parent, expectedChildren := range testCase.parentsToChildren {
            actualChildren := heap.GetChildren(heap.GetNode(0)) // TODO: fix this
            fmt.Println("actualChildren is ", actualChildren)
            if !compareChildren(expectedChildren, actualChildren) {
                t.Errorf(
                    "Received children %v for parent %d. Expected value was %v",
                    actualChildren, parent, expectedChildren,
                )
            }
        }
    }
}