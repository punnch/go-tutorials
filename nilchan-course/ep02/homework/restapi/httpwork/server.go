package httpwork

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	httpHandlers *HTTPHandlers
}

func NewHttpServer(httpHandlers *HTTPHandlers) *HTTPServer {
	return &HTTPServer{
		httpHandlers: httpHandlers,
	}
}

func (s *HTTPServer) Start() error {
	router := mux.NewRouter()

	router.Path("/books").Methods("POST").HandlerFunc(s.httpHandlers.HandleCreateBook)
	router.Path("/books/{title}").Methods("PATCH").HandlerFunc(s.httpHandlers.HandleReadBook)
	router.Path("/books/{title}").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetBook)
	router.Path("/books").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetBooks)
	router.Path("/books/{title}").Methods("DELETE").HandlerFunc(s.httpHandlers.HandleDeleteBook)

	if err := http.ListenAndServe(":9091", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
	return nil
}
