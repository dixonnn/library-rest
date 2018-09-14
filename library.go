package main

import "fmt"

var idCounter int
var books Books

// Init function is called when the program is first executed.
// Adds two books to the Book slice so data is present when
// the Index GET is first called.
func init() {
	LibraryAddBook(
		Book{
			Id:     1,
			Title:  "The Great Gatsby",
			Author: "F. Scott Fitzgerald",
			Year:   1925,
		})
	LibraryAddBook(
		Book{
			Id:     2,
			Title:  "The Catcher in the Rye",
			Author: "J.D. Salinger",
			Year:   1951,
		})
}

// Retrieves book data from the library given a book's Id.
func LibraryFindBook(id int) Book {
	for _, b := range books {
		if b.Id == id {
			return b
		}
	}
	// else, no Book with that id was found, return empty Book
	return Book{}
}

// Adds entry for new book in the library given its Book data.
// Also auto-increments the Book.Id object for each entry.
func LibraryAddBook(newBook Book) Book {
	fmt.Println("Adding new book...")
	fmt.Println(newBook)
	idCounter += 1
	newBook.Id = idCounter
	books = append(books, newBook)
	return newBook
}

// Removes an entry in the library based on the book's Id.
func LibraryRemoveBook(id int) error {
	for i, b := range books {
		if b.Id == id {
			books = append(books[:i], books[i+1:]...)
			return nil
		}
	}
	// Else return error
	return fmt.Errorf("Book with Id of %d not present", id)
}
