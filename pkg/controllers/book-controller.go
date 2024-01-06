package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bijan/go-bookstore/pkg/models"
	"github.com/bijan/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book


func GetBook(w http.ResponseWriter, r *http.Request){
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func GetBookById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	ID, err := strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book, _ := models.GetBookById(ID)
	res,_ := json.Marshal(book)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	book := &models.Book{}
	utils.ParseBody(r, book)
	b := book.CreateBook()
	res,_ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	res,_ := json.Marshal(book)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func UpdateBook(w http.ResponseWriter, r *http.Request){
	book := &models.Book{}
	utils.ParseBody(r,book)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID,err := strconv.ParseInt(bookId,0,0) 
		if err != nil {
		fmt.Println("error while parsing")
	}
	existingBook, db := models.GetBookById(ID)
	if existingBook.Name != "" {
		existingBook.Name = book.Name
	}
	if existingBook.Publication != "" {
		existingBook.Publication = book.Publication
	}
	if existingBook.Author != "" {
		existingBook.Author = book.Author
	}
	db.Save(existingBook)
	res,_ := json.Marshal(existingBook)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}