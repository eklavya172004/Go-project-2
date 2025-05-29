package controllers

import(
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/eklavya/go-bookstore/pkg/models"
	"github.com/gorilla/mux"
	"strconv"
	"github.com/eklavya/go-bookstore/pkg/utils"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter,r *http.Request){
	newBooks := models.GetallBooks()
	res,_ := json.Marshal(newBooks)
	w.Header().Set("Content-Type","application/json")// Set the response header to indicate JSON content
	//pkflication/json is a common MIME type used to indicate that the content being sent or received is in JSON format.
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBooksById(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r) // Extracts the variables from the URL
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("error while parsing bookId")
	}

	bookDetails,_ := models.GetBooksById(ID)

	res,_ := json.Marshal(bookDetails)

	w.Header().Set("Content-Type","application/json") // Set the response header to indicate JSON content
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBooks(w http.ResponseWriter,r *http.Request){
	CreateBooks := &models.Book{}
	utils.Parsebody(r,CreateBooks)

	b:= CreateBooks.CreateBook()

	res,_:=json.Marshal(b)
	// w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBooks(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID,error := strconv.ParseInt(bookId,0,0)
	if error != nil {
		fmt.Println("error while parsing bookId")
	}

	book := models.DeleteBookById(ID)

	res , _ := json.Marshal(book)

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBooks(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.Parsebody(r, updateBook)

	vars := mux.Vars(r)
	bookId := vars["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	bookdetails, db := models.GetBooksById(ID)
	if db.Error != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// Update fields only if provided
	if updateBook.Name != "" {
		bookdetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookdetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookdetails.Publication = updateBook.Publication
	}

	db.Save(bookdetails)

	res, _ := json.Marshal(bookdetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
