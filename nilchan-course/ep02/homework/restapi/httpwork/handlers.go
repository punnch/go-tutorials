package httpwork

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restapi/library"
	"strconv"

	"github.com/gorilla/mux"
)

type HTTPHandlers struct {
	bookList *library.List
}

func NewHttpHandlers(bookList *library.List) *HTTPHandlers {
	return &HTTPHandlers{
		bookList: bookList,
	}
}

/*
pattern: /books
method: POST
info: JSON

succeed:
  - status code: 201 Created
  - response body: JSON represented book

fail:
  - status code: 400, 409, 500...
  - response body: JSON represented error + time
*/
func (h *HTTPHandlers) HandleCreateBook(w http.ResponseWriter, r *http.Request) {
	var bookDTO BookDTO
	err := json.NewDecoder(r.Body).Decode(&bookDTO)
	if err != nil {
		ErrJSON(w, err, http.StatusBadRequest)
		return
	}

	book := library.NewBook(bookDTO.title, bookDTO.author, bookDTO.pages)

	if err := h.bookList.AddBook(book); err != nil {
		ErrCompareJSON(w, err, library.ErrBookAlreadyExist, http.StatusConflict)
		return
	}

	b := ToJSON(book)

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response:", err)
		return
	}
}

/*
pattern: /books/{title}
method: PATCH
info: pattern + JSON

succeed:
  - status code: 200 OK
  - response body: JSON represented changed book

fail:
  - status code: 400, 404, 500...
  - response body: JSON represented error + time
*/
func (h *HTTPHandlers) HandleReadBook(w http.ResponseWriter, r *http.Request) {
	var bookReadDTO BookReadDTO
	if err := json.NewDecoder(r.Body).Decode(&bookReadDTO); err != nil {
		ErrJSON(w, err, http.StatusBadRequest)
		return
	}

	title := mux.Vars(r)["title"]

	book, err := h.bookList.ReadBook(title, bookReadDTO.isRead)
	if err != nil {
		ErrCompareJSON(w, err, library.ErrBookNotFound, http.StatusNotFound)
		return
	}

	b := ToJSON(book)

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response:", err)
		return
	}
}

/*
pattern: /books/{title}
method: GET
info: pattern

succeed:
  - status code: 200 OK
  - response body: JSON represented received book

fail:
  - status code: 404, 500...
  - response body: JSON represented error + time
*/
func (h *HTTPHandlers) HandleGetBook(w http.ResponseWriter, r *http.Request) {

}

/*
pattern: /books?author=author_name&read=bool
method: GET
info: query params

succeed:
  - status code: 200 OK
  - response body: JSON represented received books

fail:
  - status code: 400, 500...
  - response body: JSON represented error + time
*/
func (h *HTTPHandlers) HandleGetBooks(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")
	readStr := r.URL.Query().Get("read")

	var read *bool
	if readStr != "" {
		readBool, err := strconv.ParseBool(readStr)
		if err != nil {
			ErrJSON(w, err, http.StatusBadRequest)
			return
		}
		read = &readBool
	}

	books := h.bookList.GetBooks(author, read)

	b := ToJSON(books)

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response:", err)
		return
	}
}

/*
pattern: /books/{title}
method: DELETE
info: pattern

succeed:
  - status code: 204 No content
  - response body: -

fail:
  - status code: 404, 500...
  - response body: JSON represented error + time
*/
func (h *HTTPHandlers) HandleDeleteBook(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]

	if err := h.bookList.DeleteBook(title); err != nil {
		ErrCompareJSON(w, err, library.ErrBookNotFound, http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
