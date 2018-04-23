package sorting

// O(nlogn) running time
func MergeSort(input []int) []int {
    inputLength := len(input)
    if inputLength <= 1 {
        return input
    }

    pivot := inputLength/2

    leftSlice := make([]int, pivot)
    copy(leftSlice, input[:pivot])
    rightSlice := make([]int, inputLength - pivot)
    copy(rightSlice, input[pivot:])

    leftSlice, rightSlice = MergeSort(leftSlice), MergeSort(rightSlice)
    return Merge(leftSlice, rightSlice)
}

// O(n) running time
func Merge(leftSlice []int, rightSlice []int) []int {
    if (leftSlice == nil && rightSlice == nil){
        return []int{}
    }
    if leftSlice == nil {
        return rightSlice
    }
    if rightSlice == nil {
        return leftSlice
    }

    leftIndex := 0
    rightIndex := 0

    leftSliceLength := len(leftSlice)
    rightSliceLength := len(rightSlice)

    mergedSlice := []int{}

    for {
        leftItem := leftSlice[leftIndex]
        rightItem := rightSlice[rightIndex]
        if leftItem <= rightItem {
            mergedSlice = append(mergedSlice, leftItem)
            leftIndex += 1
        } else {
            mergedSlice = append(mergedSlice, rightItem)
            rightIndex += 1
        }

        if leftIndex == (leftSliceLength){
            mergedSlice = append(mergedSlice, rightSlice[rightIndex:]...)
            break
        }
        if rightIndex == (rightSliceLength){
            mergedSlice = append(mergedSlice, leftSlice[leftIndex:]...)
            break
        }
    }

    return mergedSlice
}