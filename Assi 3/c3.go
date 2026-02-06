//demonstrate working of slices(like append ,remove,copy etc)
package main

import (
    "fmt"
)

func main() {
    slice1 := []int{1, 2, 3}
    fmt.Println("Original Slice:", slice1)
    
    // Append elements
    slice1 = append(slice1, 4, 5)
    fmt.Println("After Append:", slice1)

    // Remove element at index 2
    indexToRemove := 2
    slice1 = append(slice1[:indexToRemove], slice1[indexToRemove+1:]...)
    fmt.Println("After Removal:", slice1)
    
    // Copy slice
    slice2 := make([]int, len(slice1))
    copy(slice2, slice1)
    fmt.Println("Copied Slice:", slice2)
}