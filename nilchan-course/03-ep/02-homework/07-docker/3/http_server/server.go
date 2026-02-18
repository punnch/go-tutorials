package http_server

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	handlers *Handlers
}

func NewServer(handlers *Handlers) *Server {
	return &Server{
		handlers: handlers,
	}
}

func (s *Server) Start() error {
	router := mux.NewRouter()

	router.HandleFunc("/employees", s.handlers.CreateEmployee).Methods("POST")
	router.HandleFunc("/employees", s.handlers.GetEmployees).Methods("GET")
	router.HandleFunc("/employees/{id}", s.handlers.DeleteEmployee).Methods("DELETE")

	err := http.ListenAndServe(":8081", router)

	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}

	return err
}
