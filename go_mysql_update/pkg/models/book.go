package models

import (
	"github.com/jinzhu/gorm"
	"go_mysql/pkg/config"
	"net/http"
)

var db *gorm.DB

type Page struct {
	gorm.Model
	Number int    `json:"number"`
	Text   string `json:"text"`
	BookID uint   `json:"book_id"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{}, &Page{})
}

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Pages       []Page `gorm:"foreignkey:BookID" json:"pages"`
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Preload("Pages").Find(&books)
	return books
}

func GetById(id int64) (*Book, *gorm.DB) {
	book := &Book{}
	updateDB := db.Preload("Pages").First(book, id)
	return book, updateDB
}

func DeleteBook(id int64) int {
	book := &Book{}

	result := db.Where("ID=?", id).First(book)
	if result.Error != nil {
		return http.StatusNotFound
	}
	db.Preload("Pages").First(book, id)
	if db.Error != nil {
		return http.StatusNotFound
	}

	db.Where("book_id = ?", id).Delete(&Page{})
	if db.Error != nil {
		return http.StatusNotFound
	}

	db.Delete(book)

	if db.Error != nil {
		return http.StatusNotFound
	}

	return http.StatusOK
}
