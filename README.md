# Simple Library REST API in Go

This repo contains a simple REST implementation in Go that can view, add, and delete books from a library. This is my first project written in Go, and the code was based on [this tutorial by Corey Lanou][cl1].

## Build & Run
You can build the project to run as an execution file with:
```sh
go build
```

Otherwise, run the files manually with:

```sh
go run book.go handlers.go library.go main.go router.go routes.go
```

## Interact
By default, the server runs on port 8080. To ensure it's running correctly, two books in JSON should appear when you visit http://localhost:8080/.

#### Fetch All Books
To see all books in the library, as mentioned above, visit http://localhost:8080/. 

#### Lookup Certain Book
If you know the Id value of a book in the library, you can look up that book's information by navigating to http://localhost:8080/lookup/{bookID}, where bookID is that book's Id. Off the bat, this should work with Id values 1 and 2.

#### Add New Book
By providing Title, Author, and Year data, you can add books to the library. I used curl to accomplish this. For example, the command
```sh
curl -H "Content-Type: application/json" -d '{"title":"Great Expectations", "author":"Charles Dickens", "year":1861}' http://localhost:8080/add
```
will add an entry for Great Expectations to the library. This can be checked by navigating to the complete contents of the library again in a browser. 

#### Remove Book
To remove a book from the library, simply navigate to http://localhost:8080/remove/{bookID} where bookID is equal to the ID value of the book you would like to remove.

[cl1]: <https://thenewstack.io/make-a-restful-json-api-go/>
