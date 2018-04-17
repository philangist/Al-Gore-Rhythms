package sorting

type insertionFunction func(input []int, item int) []int

// Insertion sort: Worst case O(n**2) steps, O(n**2) comparisons
func InsertionSort(input []int, insert insertionFunction) []int {
    if insert == nil {
        insert = linearSearchInsert
    }

    output := []int{}

    for i := 0; i < len(input); i++ {
        item := input[i] 
        output = insert(output, item)
    }

    return output
}

func linearSearchInsert(input []int, item int) []int {
    inputLength := len(input)

    if inputLength == 0 {
        return []int{item}
    }

    if inputLength == 1 {
        if item < input[0] {
            return []int{item, input[0]}
        }
        return append(input, item)
    }

    var output []int
    if item <= input[0] {
        output = append([]int{item}, input...)
        return output
    }

    if item >= input[inputLength - 1] {
        output = append(input, item)
        return output        
    }

    for i := 0; i < inputLength; i++ {
        if (item <= input[i]){          
            output = append(output, input[0:i]...)
            output = append(output, item)

            // don't assign rightSlice to input[i:] directly because
            // changes to the underlying array will surface
            rightSlice := make([]int, 0)
            rightSlice = append(rightSlice, input[i:]...)
            output = append(output, rightSlice...)
            return output
        }
    }
    return input
}

// Insertion sort still takes O(n**2) steps but now we can have O(nlogn) comparisons
func binarySearchInsert(input []int, item int) []int {
    inputLength := len(input)

    if inputLength == 0 {
        return []int{item}
    }

    if inputLength == 1 {
        if item < input[0] {
            return []int{item, input[0]}
        }
        return append(input, item)
    }

    pivotIndex := inputLength/2
    pivot := input[pivotIndex]
    leftSlice := input[:pivotIndex]

    // don't assign rightSlice to input[pivotIndex:] directly because
    // changes to the underlying array will surface
    rightSlice := make([]int, 0)
    rightSlice = append(rightSlice, input[pivotIndex:]...)

    if item <= pivot {
        leftSlice = binarySearchInsert(leftSlice, item)
    }else{
        rightSlice = binarySearchInsert(rightSlice, item)
    }
    output := append(leftSlice, rightSlice...)

    return output
}
