package sorting

import (
    "fmt"
    "reflect"
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
        {
            tag: "Case 6",
            heapSlice: []int{100, 19, 36, 17, 12, 25, 5, 9, 15, 6, 11, 13, 8, 1, 4},
            valid: true,
        },
        {
            tag: "Case 7",
            heapSlice: []int{100, 19, 36, 17, 12, 25, 5, 9, 15, 6, 11, 13, 108, 1, 4},
            valid: false,
        },
        {
            tag: "Case 8",
            heapSlice: []int{36, 19, 25, 7,17, 1, 2, 3},
            valid: true,
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
type heapPopRootTestCases struct {
    tag string
    initialHeap []int
    reorderedHeap []int
    heapRootValue int
}

func TestPopHeapRoot(t *testing.T){
    fmt.Println("Running TestPopHeapRoot...")

    testCases := []heapPopRootTestCases{
        {
            tag: "Case 1",
            initialHeap: []int{100, 19, 36, 17, 12, 25, 5, 9, 15, 6, 11, 13, 8, 1, 4},
            reorderedHeap: []int{36, 19, 25, 15, 17, 13, 9, 12, 6, 11, 8, 5, 1, 4},
            heapRootValue: 100,
        },
        {
            tag: "Case 32",
            initialHeap: []int{87, 19, 36, 17, 3, 25, 1, 2, 7},
            reorderedHeap: []int{36, 19, 25, 7, 17, 1, 2, 3},
            heapRootValue: 87,
        },
        {
            tag: "Case 3",
            initialHeap: []int{5,4,3,2,1},
            reorderedHeap: []int{4,3,2,1},
            heapRootValue: 5,
        },
        {
            tag: "Case 4",
            initialHeap: []int{4,3,2,1},
            reorderedHeap: []int{3,2,1},
            heapRootValue: 4,
        },
        {
            tag: "Case 5",
            initialHeap: []int{3,2,1},
            reorderedHeap: []int{2,1},
            heapRootValue: 3,
        },
        {
            tag: "Case 6",
            initialHeap: []int{2,1},
            reorderedHeap: []int{1},
            heapRootValue: 2,
        },
        {
            tag: "Case 7",
            initialHeap: []int{1},
            reorderedHeap: []int{},
            heapRootValue: 1,
        },
        {
            tag: "Case 8",
            initialHeap: []int{},
            reorderedHeap: []int{},
            heapRootValue: NIL_VALUE,
        },
    }

    for _, testCase := range testCases {
        fmt.Println("Test pop root of heap: ", testCase.initialHeap)

        heap := NewHeap(testCase.initialHeap)
        heapRoot, err := heap.PopRoot()

        if (err != nil) {
            t.Errorf("%s: Pop heap root failed. Received unexpected error '%s'", testCase.tag, err)
        }
        if (reflect.DeepEqual(heap.Values, testCase.reorderedHeap) == false){
            t.Errorf(
                "%s: Heap reorder failed.\nExpected values: %v\nReceived values:%v",
                testCase.tag, testCase.reorderedHeap, heap.Values,
            )
        }
        if (heapRoot.Value != NIL_VALUE && (heapRoot.Value != testCase.heapRootValue)){
            t.Errorf(
                "%s: returned heapRoot value is incorrect. Expected: %d Received: %d",
                testCase.tag, testCase.heapRootValue, heapRoot.Value,
            )
        }
    }
}

func TestHeapInsert(t *testing.T){
    fmt.Println("Running TestHeapInsert...")
    inputValues := []int{1, 2, 3, 7, 19, 25, 36, 17}
    outputValues := []int{36, 25, 7, 19, 1, 3, 2, 17}

    heap := NewHeap([]int{})
    for _, inputValue := range inputValues {
        heap.InsertValue(inputValue)
    }
    if (reflect.DeepEqual(heap.Values, outputValues) == false){
        t.Errorf(
            "Insertion of heap values did not generate expected order\nExpected Values: %v\nActual values:%v ",
            outputValues, heap.Values,
        )
    }
}