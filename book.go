package main

// Book is a struct that stores information about books.
// This REST API can crudely store Book structs in a
type Book struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

// Books is a custom type used by library.go to crudely
// store Book structs. This REST API performs operations
// on a single Books slice.
type Books []Book
