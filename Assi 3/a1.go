//wap in go lang. to find the largest and smallest number in an array.

package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter the number of elements in the array: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	fmt.Println("Enter the elements of the array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	largest := arr[0]
	smallest := arr[0]

	for i := 1; i < n; i++ {
		if arr[i] > largest {
			largest = arr[i]
		}
		if arr[i] < smallest {
			smallest = arr[i]
		}
	}

	fmt.Printf("Largest number: %d\n", largest)
	fmt.Printf("Smallest number: %d\n", smallest)
}	
