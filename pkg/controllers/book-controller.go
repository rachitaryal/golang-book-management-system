package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rachitaryal/golang-book-management-system/pkg/models"
	"github.com/rachitaryal/golang-book-management-system/pkg/utils"
)

func GetBooks(res http.ResponseWriter, req *http.Request){
	books := models.GetAllBooks()
	data, _ := json.Marshal(books)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(data)
}

func GetBookById(res http.ResponseWriter, req *http.Request){
	vars := mux.Vars(req)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error While Parsing bookId ", bookId)
	}
	book, _ := models.GetBookById(ID)
	data, _ := json.Marshal(book)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(data)

}

func CreateBook(res http.ResponseWriter, req *http.Request){
	BookInstance := &models.Book{}
	utils.ParseBody(req, BookInstance)
	book := BookInstance.CreateBook()
	data, _ := json.Marshal(book)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(data)

}

func DeleteBook(res http.ResponseWriter, req *http.Request){
	vars := mux.Vars(req)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error While Parsing bookId ", bookId)
	}
	book := models.DeleteBookById(ID)
	data, _ := json.Marshal(book)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(data)

}

func UpdateBook(res http.ResponseWriter, req *http.Request){
	updateBook := &models.Book{}
	utils.ParseBody(req, updateBook)
	vars := mux.Vars(req)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error While Parsing bookId ", bookId)
	}
	book, db := models.GetBookById(ID)
	if updateBook.Name != ""{
		book.Name = updateBook.Name
	}
	if updateBook.Author != ""{
		book.Author = updateBook.Author
	}
	if updateBook.Publication != ""{
		book.Publication = updateBook.Publication
	}
	db.Save(&book)
	data, _ := json.Marshal(book)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(data)


}