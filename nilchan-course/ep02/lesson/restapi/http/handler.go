package http

// main goal: what should I send or get from client when...

import (
	"net/http"
	"restapi/todo"
)

type HTTPHandlers struct {
	todoList *todo.List
}

func NewHTTPHandlers(todoList *todo.List) *HTTPHandlers {
	return &HTTPHandlers{
		todoList: todoList,
	}
}

/*
pattern: /tasks
method: POST
info: JSON in HTTP request body

succeed:
  - status code: 201 Created
  - response body: JSON represented created task

failed:
  - status code: 400, 409, 500...
  - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandleCreateTask(w http.ResponseWriter, r *http.Request) {

}

/*
pattern: /tasks/{title}
method: GET
info: pattern

succeed:
  - status code: 200 OK
  - response body: JSON represented task

failed:
  - status code: 400, 404, 500...
  - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandleGetTask(w http.ResponseWriter, r *http.Request) {

}

/*
pattern: /tasks
method: GET
info: -

succeed:
  - status code: 200 OK
  - response body: JSON represented tasks

failed:
  - status code: 400, 500...
  - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandleGetTasks(w http.ResponseWriter, r *http.Request) {

}

/*
pattern: /tasks?completed=false
method: GET
info: query params

succeed:
  - status code: 200 OK
  - response body: JSON represented uncompleted tasks

failed:
  - status code: 400, 500...
  - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandleGetUncompletedTasks(w http.ResponseWriter, r *http.Request) {

}

/*
pattern: /tasks/{title}
method: PATCH
info: pattern + JSON

succeed:
  - status code: 200 OK
  - response body: JSON represented changed task

failed:
  - status code: 400, 404, 500...
  - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandleCompleteTask(w http.ResponseWriter, r *http.Request) {

}

/*
pattern: /tasks/{title}
method: DELETE
info: pattern

succeed:
  - status code: 204 No content
  - response body: -

failed:
  - status code: 400, 404, 500...
  - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandleDeleteTask(w http.ResponseWriter, r *http.Request) {

}
