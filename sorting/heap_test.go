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
    tag       string
    heapSlice []int
    parentsToChildren map[int][]int //parent index -> (leftChildIndex, rightChildIndex)
}

func TestHeapChildren(t *testing.T){
    fmt.Println("Running TestHeapChildren...")
    testCases := []heapChildrenTestCases{
        {
            tag: "Case 1",
            heapSlice: []int{4, 9, 7},
            parentsToChildren: map[int][]int{
                4: []int{9, 7},
            },
        },
        {
            tag: "Case ",
            heapSlice: []int{16, -1, -1},
            parentsToChildren: map[int][]int{
                16: []int{NIL_VALUE, NIL_VALUE},
            },
        },
    }

    compare := func(expectedChildren []int, actualChildren []*HeapNode) bool {
        return ((expectedChildren[0] == actualChildren[0].Value) &&
                (expectedChildren[1] == actualChildren[1].Value))
    }

    for _, testCase := range testCases {
        heap := NewHeap(testCase.heapSlice)

        for parent, expectedChildren := range testCase.parentsToChildren {
            fmt.Println("Testing parent child relationship for heap: ", parent, expectedChildren)
            actualChildren := heap.GetChildren(heap.GetNode(0)) // TODO: fix this
            if !compare(expectedChildren, actualChildren) {
                t.Errorf(
                    "Received children %v for parent %d. Expected value was %v",
                    actualChildren, parent, expectedChildren,
                )
            }
        }
    }
}

type heapInvariantTestCases struct {
    tag string
    heapSlice []int
    valid bool
}

func TestHeapInvariant (t *testing.T){
    fmt.Println("Running TestHeapInvariant...")

    testCases := []heapInvariantTestCases{
        {
            tag: "Case 1",
            heapSlice: []int{1, 2, 3},
            valid: false,
        },
        {
            tag: "Case 2",
            heapSlice: []int{6, 2, 7},
            valid: false,
        },
        {
            tag: "Case 3",
            heapSlice: []int{6, 2, 3},
            valid: true,
        },
        {
            tag: "Case 4",
            heapSlice: []int{100, 19, 36, 17, 3, 25, 1, 2, 7},
            valid: true,
        },
        {
            tag: "Case 5",
            heapSlice: []int{100, 19, 36, 17, 3, 25, 1, 2, 7},
            valid: false,
        },
    }

    for _, testCase := range testCases {
        fmt.Println("Test heap invariant of heap: ", testCase.heapSlice)

        heap := NewHeap(testCase.heapSlice)
        err := heap.assertInvariant()

        if (err != nil) && testCase.valid == true {
            t.Errorf("%s: Heap invariant does not hold. Received unexpected error '%s'", testCase.tag, err)
        }
        if (err != nil) && testCase.valid == false {
            fmt.Println(testCase.tag, ": Heap invariant holds. Expected validation state(failure) was met")
        }
        if (err == nil) && testCase.valid == false {
            fmt.Errorf("%s: Heap invariant does not hold. Expected validation state(success) was not met", testCase.tag)
        }
        if (err == nil) && testCase.valid == true {
            fmt.Println(testCase.tag, ": Heap invariant holds. Expected validation state(success) was met")
        }
    }
}