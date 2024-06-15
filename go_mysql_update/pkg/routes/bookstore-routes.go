package routes

import (
	"github.com/gorilla/mux"
	"go_mysql/pkg/controllers"
	"net/http"
)

var BookStoreRoutes = func(r *mux.Router) {
	r.HandleFunc("/book/", controllers.CreateBook).Methods(http.MethodPost)
	r.HandleFunc("/book/", controllers.GetBooks).Methods(http.MethodGet)
	r.HandleFunc("/book/{bookid}", controllers.GetBookById).Methods(http.MethodGet)
	r.HandleFunc("/book/{bookid}", controllers.DeleteBook).Methods(http.MethodDelete)
	r.HandleFunc("/book/{bookid}", controllers.UpdateBook).Methods(http.MethodPut)
}
