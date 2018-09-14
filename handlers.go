package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Index is the handler function for the index page of the library server.
// Accessing http://localhost:8080/ will send a GET request for the current
// books in the library.
func Index(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(books); err != nil {
		panic(err)
	}
}

// LookupBook is a function that requests a Book structure from the
// library based on a known ID number. It is called when accessing
// the path http://localhost:8080/lookup/{bookId}.
func LookupBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var bookId int
	var err error

	// First, retrieve the bookId value from the request's arg.
	if bookId, err = strconv.Atoi(vars["bookId"]); err != nil {
		panic(err)
	}

	// Search the library for the book based on its Id.
	book := LibraryFindBook(bookId)
	if book.Id > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(book); err != nil {
			panic(err)
		}
		return
	}

	// If book not present, return null.
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode("Book not found, nothing happened."); err != nil {
		panic(err)
	}

}

// Allows the addition of a book to the library given its data.
// Send POST request to http://localhost:8080/add.
func AddBook(w http.ResponseWriter, r *http.Request) {
	var book Book

	// Ensure body is not unreasonably large
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	// Must Close() body when we're finished with it
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	// Map values in Body to Book struct.
	if err := json.Unmarshal(body, &book); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	// As long as the values are mapped correctly, add book to Library
	b := LibraryAddBook(book)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(b); err != nil {
		panic(err)
	}
}

// Removes book entry from library based on Book.Id value.
// Send POST request to http://localhost:8080/remove/{bookId}
func RemoveBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var bookId int
	var err error

	// First, retrieve the bookId value from the request's arg.
	if bookId, err = strconv.Atoi(vars["bookId"]); err != nil {
		panic(err)
	}

	// Call LibraryRemoveBook to operate on library slice
	book := LibraryRemoveBook(bookId)
	if book == nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode("Book removed from library."); err != nil {
			panic(err)
		}
		return
	}

	// If book not present, return null
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode("Book not found, nothing happened."); err != nil {
		panic(err)
	}
}
