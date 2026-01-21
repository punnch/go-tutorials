package http

// main goal: what should I send or get from client when... (client-server interaction)

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"restapi/todo"
	"time"
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
	// 1. get json
	var taskDTO TaskDTO
	if err := json.NewDecoder(r.Body).Decode(&taskDTO); err != nil {
		errorDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errorDTO.ToString(), http.StatusBadRequest)
		return
	}
	// 2. validate task
	if err := taskDTO.ValidateForCreate(); err != nil {
		errorDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errorDTO.ToString(), http.StatusBadRequest)
		return
	}

	// 3. add task to list
	todoTask := todo.NewTask(taskDTO.Title, taskDTO.Description)
	if err := h.todoList.AddTask(todoTask); err != nil {
		errorDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		if errors.Is(err, todo.ErrTaskAlreadyExists) {
			http.Error(w, errorDTO.ToString(), http.StatusConflict)
		} else {
			http.Error(w, errorDTO.ToString(), http.StatusInternalServerError)
		}
		return
	}

	// 4. send task to the server
	b, err := json.MarshalIndent(taskDTO, "", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response:", err)
	}
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
func (h *HTTPHandlers) HandleGetAllTasks(w http.ResponseWriter, r *http.Request) {

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
func (h *HTTPHandlers) HandleGetAllNotCompletedTasks(w http.ResponseWriter, r *http.Request) {

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
