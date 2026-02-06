//accept the book details such as book id, book name, author name from the user and display the book details.
package main

import "fmt"

type Book struct {
	id     int
	name   string
	author string
}

func main() {
	var book Book

	fmt.Print("Enter Book ID: ")
	fmt.Scan(&book.id)

	fmt.Print("Enter Book Name: ")
	fmt.Scan(&book.name)

	fmt.Print("Enter Author Name: ")
	fmt.Scan(&book.author)

	fmt.Println("\nBook Details:")
	fmt.Printf("Book ID: %d\n", book.id)
	fmt.Printf("Book Name: %s\n", book.name)
	fmt.Printf("Author Name: %s\n", book.author)
}	 