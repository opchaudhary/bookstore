// handlers/book_handler.go
package handlers

import (
    "encoding/json"
    "net/http"

    "library_api/models"
)

// Error response
func errorResponse(w http.ResponseWriter, message string, code int) {
    w.WriteHeader(code)
    json.NewEncoder(w).Encode(map[string]string{"error": message})
}

var books []models.Book

// Add a book
func AddBook(w http.ResponseWriter, r *http.Request) {
    var newBook models.Book
    err := json.NewDecoder(r.Body).Decode(&newBook)
    if err != nil {
        errorResponse(w, "Invalid request body", http.StatusBadRequest)
        return
    }
    books = append(books, newBook)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newBook)
}

// Get all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(books)
}

// Get a single book by ID
func GetBook(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    for _, book := range books {
        if book.ID == id {
            json.NewEncoder(w).Encode(book)
            return
        }
    }
    errorResponse(w, "Book not found", http.StatusNotFound)
}

// Update a book by ID
func UpdateBook(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    var updatedBook models.Book
    err := json.NewDecoder(r.Body).Decode(&updatedBook)
    if err != nil {
        errorResponse(w, "Invalid request body", http.StatusBadRequest)
        return
    }
    for i, book := range books {
        if book.ID == id {
            books[i] = updatedBook
            json.NewEncoder(w).Encode(updatedBook)
            return
        }
    }
    errorResponse(w, "Book not found", http.StatusNotFound)
}

// Delete a book by ID
func DeleteBook(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    for i, book := range books {
        if book.ID == id {
            books = append(books[:i], books[i+1:]...)
            w.WriteHeader(http.StatusNoContent)
            return
        }
    }
    errorResponse(w, "Book not found", http.StatusNotFound)
}
