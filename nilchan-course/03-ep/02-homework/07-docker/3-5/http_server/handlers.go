package http_server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"company/company"
	"company/http_server/dto"
)

type Handlers struct {
	company *company.Company
}

func NewHandlers(company *company.Company) *Handlers {
	return &Handlers{
		company: company,
	}
}

func (h *Handlers) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var employeeDTO dto.EployeeDTO
	if err := json.NewDecoder(r.Body).Decode(&employeeDTO); err != nil {
		dto.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	employee, err := h.company.CreateEmployee(r.Context(), employeeDTO.FullName, employeeDTO.Position)
	if err != nil {
		dto.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	b := dto.ToJSON(employee)

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write response body:", err)
		return
	}
}

func (h *Handlers) GetEmployees(w http.ResponseWriter, r *http.Request) {
	employees, err := h.company.GetEmployees(r.Context())
	if err != nil {
		dto.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	b := dto.ToJSON(employees)

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write response body:", err)
		return
	}
}

func (h *Handlers) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		dto.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := h.company.DeleteEmployee(r.Context(), id); err != nil {
		dto.ErrorCompareJSON(w, err, company.ErrNotFound, http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
