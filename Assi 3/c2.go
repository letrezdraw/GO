//accept n records of employee information (ename ,id,salary) using structure and display the employee details having maximum salary.
package main

import "fmt"

type Employee struct {
	id     int
	ename  string
	salary float64
}

func main() {
	var n int
	fmt.Print("Enter the number of employees: ")
	fmt.Scan(&n)

	employees := make([]Employee, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Enter details for employee %d:\n", i+1)
		fmt.Print("Employee ID: ")
		fmt.Scan(&employees[i].id)
		fmt.Print("Employee Name: ")
		fmt.Scan(&employees[i].ename)
		fmt.Print("Employee Salary: ")
		fmt.Scan(&employees[i].salary)
	}

	maxSalaryEmployee := employees[0]
	for i := 1; i < n; i++ {
		if employees[i].salary > maxSalaryEmployee.salary {
			maxSalaryEmployee = employees[i]
		}
	}

	fmt.Println("\nEmployee with maximum salary:")
	fmt.Printf("Employee ID: %d\n", maxSalaryEmployee.id)
	fmt.Printf("Employee Name: %s\n", maxSalaryEmployee.ename)
	fmt.Printf("Employee Salary: %.2f\n", maxSalaryEmployee.salary)
}	