package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go_mysql/pkg/models"
	"go_mysql/pkg/utils"
	"log"
	"net/http"
	"strconv"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(r, book)
	b := book.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	allbooks := models.GetAllBooks()
	res, err := json.Marshal(allbooks)
	if err != nil {
		http.Error(w, "Ошибка при маршалинге данных", http.StatusInternalServerError)
		log.Printf("Ошибка при маршалинге данных: %v", err)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["bookid"], 0, 0)
	if err != nil {
		panic(err)
	}
	book, _ := models.GetById(id)
	res, _ := json.Marshal(book)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["bookid"], 0, 0)
	if err != nil {
		panic(err)
	}
	b := models.DeleteBook(id)
	w.WriteHeader(b)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	parseBook := &models.Book{}
	utils.ParseBody(r, parseBook)
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["bookid"], 0, 0)
	if err != nil {
		panic(err)
	}
	dbBook, db := models.GetById(id)
	if parseBook.Name != "" {
		dbBook.Name = parseBook.Name
	}
	if parseBook.Author != "" {
		dbBook.Author = parseBook.Author
	}
	if parseBook.Description != "" {
		dbBook.Description = parseBook.Description
	}
	if parseBook.Name != "" {
		dbBook.Description = parseBook.Description
	}
	if len(parseBook.Pages) > 0 {
		mapExistedPages := make(map[int]*models.Page)
		for i := range dbBook.Pages {
			mapExistedPages[dbBook.Pages[i].Number] = &dbBook.Pages[i]
		}
		for _, page := range parseBook.Pages {
			if val, exists := mapExistedPages[page.Number]; exists {
				val.Text = page.Text
			}
		}
	}
	db.Save(dbBook)
	res, _ := json.Marshal(dbBook)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
