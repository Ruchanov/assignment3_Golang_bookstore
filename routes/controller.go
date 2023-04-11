package routes

import (
	"encoding/json"
	"fmt"
	"github.com/Ruchanov/Golang_2023/assignment3/models"
	"github.com/Ruchanov/Golang_2023/assignment3/repositories"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type Controller struct {
	repository *repositories.BookRepository
}

func NewController(db *gorm.DB) *Controller {
	return &Controller{
		repository: repositories.NewBookRepository(db),
	}
}

func (c *Controller) HandleBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		res := c.repository.GetBooks()
		//if err != nil {
		//	errorResponse(w, http.StatusInternalServerError, err)
		//	return
		//}
		response(w, http.StatusOK, res)
	} else if r.Method == "POST" {
		book := models.Book{}
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			response(w, http.StatusBadRequest, err)
			return
		}
		err = c.repository.CreateBook(&book)
		if err != nil {
			response(w, http.StatusInternalServerError, err)
			return
		}
		response(w, http.StatusCreated, nil)
	}
}

func (c *Controller) HandleBookById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		errorResponse(w, http.StatusBadRequest, err)
		return
	}
	book, err := c.repository.GetBookByID(id)
	if err != nil {
		errorResponse(w, http.StatusBadRequest, err)
		return
	}
	if r.Method == "GET" {
		response(w, http.StatusOK, book)
	} else if r.Method == "DELETE" {
		err = c.repository.DeleteBook(book)
		if err != nil {
			errorResponse(w, http.StatusInternalServerError, err)
			return
		}
		response(w, http.StatusNoContent, nil)
	} else if r.Method == "PUT" {
		var updatedBook models.Book
		err = json.NewDecoder(r.Body).Decode(&updatedBook)
		if err != nil {
			errorResponse(w, http.StatusInternalServerError, err)
			return
		}
		//var b models.Book
		_, err = c.repository.UpdateBookByID(id, &updatedBook)
		if err != nil {
			errorResponse(w, http.StatusInternalServerError, err)
			return
		}
		response(w, http.StatusCreated, nil)
	}

}

func (c *Controller) HandleSearch(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	books, err := c.repository.SearchByTitle(title)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, err)
		return
	}
	response(w, http.StatusOK, books)
}

func (c *Controller) HandleOrder(w http.ResponseWriter, r *http.Request) {
	order := mux.Vars(r)["order"]
	fmt.Println(order)
	books, err := c.repository.GetBooksSortedByCost(order)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, err)
		return
	}
	response(w, http.StatusOK, books)
}

func errorResponse(w http.ResponseWriter, code int, err error) {
	response(w, code, map[string]string{"error": err.Error()})
}

func response(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
