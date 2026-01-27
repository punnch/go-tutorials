package http

// main goal: what should I send or get from client when... (client-server interaction)

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restapi/todo"

	"github.com/gorilla/mux"
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
	var taskDTO TaskDTO
	if err := json.NewDecoder(r.Body).Decode(&taskDTO); err != nil {
		errorDTO := NewErrorDTO(err.Error())

		http.Error(w, errorDTO.ToString(), http.StatusBadRequest)
		return
	}

	if err := taskDTO.ValidateForCreate(); err != nil {
		errorDTO := NewErrorDTO(err.Error())

		http.Error(w, errorDTO.ToString(), http.StatusBadRequest)
		return
	}

	todoTask := todo.NewTask(taskDTO.Title, taskDTO.Description)
	if err := h.todoList.AddTask(todoTask); err != nil {
		errorDTO := NewErrorDTO(err.Error())

		errorDTO.CompareSendErr(w, err, todo.ErrTaskAlreadyExists, http.StatusConflict)
		return
	}

	b, err := json.MarshalIndent(todoTask, "", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response:", err)
		return
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
	title := mux.Vars(r)["title"]

	task, err := h.todoList.GetTask(title)
	if err != nil {
		errorDTO := NewErrorDTO(err.Error())

		errorDTO.CompareSendErr(w, err, todo.ErrTaskNotFound, http.StatusNotFound)
		return
	}

	b, err := json.MarshalIndent(task, "", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response:", err)
		return
	}
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
	tasks := h.todoList.GetTasks()

	b, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response:", err)
		return
	}
}

/*
pattern: /tasks?completed=true
method: GET
info: query params

succeed:
  - status code: 200 OK
  - response body: JSON represented uncompleted tasks

failed:
  - status code: 400, 500...
  - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandleGetCompletedTasks(w http.ResponseWriter, r *http.Request) {
	completedTasks := h.todoList.GetCompletedTasks()

	b, err := json.MarshalIndent(completedTasks, "", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response:", err)
		return
	}
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
	uncompletedTasks := h.todoList.GetUncompletedTasks()

	b, err := json.MarshalIndent(uncompletedTasks, "", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response:", err)
		return
	}
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
	var completeTaskDTO CompleteTaskDTO
	if err := json.NewDecoder(r.Body).Decode(&completeTaskDTO); err != nil {
		errorDTO := NewErrorDTO(err.Error())

		http.Error(w, errorDTO.ToString(), http.StatusBadRequest)
		return
	}

	title := mux.Vars(r)["title"]

	var (
		task todo.Task
		err  error
	)

	if completeTaskDTO.Completed {
		task, err = h.todoList.CompleteTask(title)
	} else {
		task, err = h.todoList.UncompleteTask(title)
	}

	if err != nil {
		errorDTO := NewErrorDTO(err.Error())

		errorDTO.CompareSendErr(w, err, todo.ErrTaskNotFound, http.StatusNotFound)
		return
	}

	b, err := json.MarshalIndent(task, "", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response:", err)
		return
	}
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
	title := mux.Vars(r)["title"]

	if err := h.todoList.DeleteTask(title); err != nil {
		errorDTO := NewErrorDTO(err.Error())

		errorDTO.CompareSendErr(w, err, todo.ErrTaskNotFound, http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
