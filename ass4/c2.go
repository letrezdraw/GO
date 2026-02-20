package main

import (
	"fmt"
	"sort"
)

type Student struct {
	RollNo     int
	Name       string
	Percentage float64
}

type ByPercentage []Student

func (a ByPercentage) Len() int           { return len(a) }
func (a ByPercentage) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPercentage) Less(i, j int) bool { return a[i].Percentage > a[j].Percentage }

func main() {
	students := []Student{
		{RollNo: 1, Name: "soham", Percentage: 85.5},
		{RollNo: 2, Name: "ajay", Percentage: 90.0},
		{RollNo: 3, Name: "raj", Percentage: 78.0},
	}

	sort.Sort(ByPercentage(students))

	fmt.Println("Students in descending order of percentage:")
	for _, student := range students {
		fmt.Printf("Roll No: %d, Name: %s, Percentage: %.2f\n", student.RollNo, student.Name, student.Percentage)
	}
}
