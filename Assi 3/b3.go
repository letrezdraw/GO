// accept n student details such as student id, name,mark1,mark2,mark3 .calculate total  and average of marks using structure.
package main

import "fmt"

type Student struct {
	id    int
	name  string
	mark1 float64
	mark2 float64
	mark3 float64
}

func main() {
	var student Student

	fmt.Print("Enter Student ID: ")
	fmt.Scan(&student.id)

	fmt.Print("Enter Student Name: ")
	fmt.Scan(&student.name)
	fmt.Print("Enter Mark 1: ")
	fmt.Scan(&student.mark1)

	fmt.Print("Enter Mark 2: ")
	fmt.Scan(&student.mark2)

	fmt.Print("Enter Mark 3: ")
	fmt.Scan(&student.mark3)

	total := student.mark1 + student.mark2 + student.mark3
	average := total / 3
	fmt.Println("\nStudent Details:")
	fmt.Printf("Student ID: %d\n", student.id)
	fmt.Printf("Student Name: %s\n", student.name)
	fmt.Printf("Total Marks: %.2f\n", total)
	fmt.Printf("Average Marks: %.2f\n", average)
}
