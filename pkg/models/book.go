package models

import (
	"github.com/jinzhu/gorm"
	"github.com/rachitaryal/golang-book-management-system/pkg/config"
)

var db *gorm.DB


type Book struct {
	gorm.Model
	Name string `json:"name"`
	Author string `json:"author"`
	Publication string`json:"publication"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func(book *Book) CreateBook() *Book{
	db.NewRecord(book)
	db.Create(&book)
	return book
}

func GetAllBooks()[]Book{
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(Id int64) (*Book, *gorm.DB){
	var book Book
	db :=db.Where("ID=?", Id).Find(&book)
	return &book, db
}

func DeleteBookById(Id int64)Book{
	var book Book
	db.Where("ID=?", Id).Delete(&book)
	return book

}