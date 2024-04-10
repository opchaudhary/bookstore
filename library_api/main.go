// main.go
package main

import (
    "net/http"

    "library_api/handlers"
)

func main() {
    http.HandleFunc("/books", handlers.GetBooks)
    http.HandleFunc("/books/add", handlers.AddBook)
    http.HandleFunc("/books/update", handlers.UpdateBook)
    http.HandleFunc("/books/delete", handlers.DeleteBook)
    
    http.ListenAndServe(":8080", nil)
}
