package http

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

// main goal: what handler should I execute in that moment if... (interaction with handlers)

type HTTPServer struct {
	httpHandlers *HTTPHandlers
}

func NewHTTPServer(httpHandlers *HTTPHandlers) *HTTPServer {
	return &HTTPServer{
		httpHandlers: httpHandlers,
	}
}

func (s *HTTPServer) StartServer() error {
	router := mux.NewRouter()

	router.Path("/tasks").Methods("POST").HandlerFunc(s.httpHandlers.HandleCreateTask)
	router.Path("/tasks/{title}").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetTask)
	router.Path("/tasks").Methods("GET").Queries("completed", "true").HandlerFunc(s.httpHandlers.HandleGetCompletedTasks)
	router.Path("/tasks").Methods("GET").Queries("completed", "false").HandlerFunc(s.httpHandlers.HandleGetUncompletedTasks)
	router.Path("/tasks").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetTasks)
	router.Path("/tasks/{title}").Methods("PATCH").HandlerFunc(s.httpHandlers.HandleCompleteTask)
	router.Path("/tasks/{title}").Methods("DELETE").HandlerFunc(s.httpHandlers.HandleDeleteTask)

	if err := http.ListenAndServe(":9091", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}

	return nil
}
