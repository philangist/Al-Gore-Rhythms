package trees

import (
    "fmt"
    "reflect"
    "testing"
)

type searchTestCases struct {
    tag       string
    inputValues *Node
    destination int
    expectedPath []int
}

func generateTestData() *Node {
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

    return node1
}

func TestDepthFirstSearch(t *testing.T){
    fmt.Println("Running TestDepthFirstSearch...")
    testCases := []searchTestCases{
        {
            tag: "Case 1",
            inputValues: generateTestData(),
            destination: 8,
            expectedPath: []int{1,7,8},
        },
        {
            tag: "Case 2",
            inputValues: generateTestData(),
            destination: 4,
            expectedPath: []int{1,2,4},
        },
        {
            tag: "Case 3",
            inputValues: generateTestData(),
            destination: 1,
            expectedPath: []int{1},
        },
        {
            tag: "Case 4",
            inputValues: generateTestData(),
            destination: 9,
            expectedPath: []int{},
        },
    }

    for _, testCase := range testCases {
        fmt.Println("Running depth first search tree structure for ", testCase.inputValues)
        actualPath := ExtractPathToRoot(DepthFirstSearch(testCase.inputValues, testCase.destination))

        if (!reflect.DeepEqual(testCase.expectedPath, actualPath)){
            t.Errorf("%s: Breadth first search failed.\nExpected path to destination %d: %v\nObserved path:%v",
                testCase.tag, testCase.destination, testCase.expectedPath, actualPath,
            )
        }
    }
}

func TestBreadthFirstSearch(t *testing.T){
    fmt.Println("Running TestBreadthFirstSearch...")
    testCases := []searchTestCases{
        {
            tag: "Case 1",
            inputValues: generateTestData(),
            destination: 8,
            expectedPath: []int{1,7,8},
        },
        {
            tag: "Case 2",
            inputValues: generateTestData(),
            destination: 4,
            expectedPath: []int{1,2,4},
        },
        {
            tag: "Case 3",
            inputValues: generateTestData(),
            destination: 1,
            expectedPath: []int{1},
        },
        {
            tag: "Case 4",
            inputValues: generateTestData(),
            destination: 9,
            expectedPath: []int{},
        },
    }

    for _, testCase := range testCases {
        fmt.Println("Running breadth first search tree structure for ", testCase.inputValues)
        actualPath := ExtractPathToRoot(BreadthFirstSearch(testCase.inputValues, testCase.destination))

        if (!reflect.DeepEqual(testCase.expectedPath, actualPath)){
            t.Errorf("%s: Breadth first search failed.\nExpected path to destination %d: %v\nObserved path:%v",
                testCase.tag, testCase.destination, testCase.expectedPath, actualPath,
            )
        }
    }
}
