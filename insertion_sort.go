package main

import "fmt"

func insertion_sort(input []int) []int {
    output := []int{}

    for i := 0; i < len(input); i++ {
        item := input[i] 
        output = insert(output, item)
        fmt.Println("insertion_sort loop output: ", output)
    }

    return output
}

func insert(input []int, item int) []int {
    input_length := len(input)

    fmt.Println("input", input)
    fmt.Println("item", item)
    if input_length == 0 {
        return []int{item}
    }

    if input_length == 1 {
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

    if item >= input[input_length - 1] {
        output = append(input, item)
        return output        
    }

    for i := 0; i < input_length; i++ {
        fmt.Println("input[",i,"] is:", input[i])
        if (item <= input[i]){          
            output = append(output, input[0:i]...)
            output = append(output, item)

             // don't assign input_right to input[i:] because
             // changes to the underlying array will surface
            input_right := make([]int, 0) 
            input_right = append(input_right, input[i:]...)
            output = append(output, input_right...)
            return output
        }
    }
    return input
}

func main() {
    input := []int{2,1,5,-1,3,4,99}
    fmt.Println("Unsorted input ", input)
    output := insertion_sort(input)
    fmt.Println("Sorted output ", output)
}