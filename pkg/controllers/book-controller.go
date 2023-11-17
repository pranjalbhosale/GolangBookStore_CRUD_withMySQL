package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pranjalbhosale/Go_BookStore/pkg/models"
	"github.com/pranjalbhosale/Go_BookStore/pkg/utils"
)

var newBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {

	// Call the GetAllBooks function to retrieve all books
	newBook := models.GetAllBooks()
	// Marshal the books slice into JSON
	res, err := json.Marshal(newBook)
	if err != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}
	// Set the response headers
	w.Header().Add("Content-Type", "pkglication/json")
	// Write the JSON response
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	// Extract the book ID from the request parameters using Gorilla Mux.
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		http.Error(w, "Error parsing book ID", http.StatusBadRequest)
		return
	}

	// Retrieve book details from the models package using the parsed book ID.
	bookDetails, _ := models.GetBookById(ID)
	// Marshal the book details into JSON format.
	res, _ := json.Marshal(bookDetails)
	w.Header().Add("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	// Create an instance of the models.Book struct to store the data from the request.
	newBook := &models.Book{}
	// Use the ParseBody utility function to parse the JSON request body and populate the newBook instance.
	utils.ParseBody(r, newBook)
	// Call the CreateBook method on the newBook instance to persist the book in the database.
	b := newBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		http.Error(w, "Error parsing book ID", http.StatusBadRequest)
		return
	}

	deleteBook := models.DeleteBook(ID)
	res, _ := json.Marshal(deleteBook)
	w.Header().Add("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		http.Error(w, "Error parsing book ID", http.StatusBadRequest)
		return
	}
	bookDetails, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}

	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}

	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)

	res, _ := json.Marshal(bookDetails)
	w.Header().Add("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
