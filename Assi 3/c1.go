//accept two matrices and display its multiplication
package main

import "fmt"

func main() {
	var row1, col1, row2, col2 int
	
	fmt.Print("Enter number of rows and columns for first matrix: ")
	fmt.Scan(&row1, &col1)
	fmt.Print("Enter number of rows and columns for second matrix: ")
	fmt.Scan(&row2, &col2)
	
	if col1 != row2 {
		fmt.Println("Matrix multiplication not possible with the given dimensions.")
		return
	}

}