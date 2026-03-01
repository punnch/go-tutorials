package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"logging/logging"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Handler struct {
	logger *zap.Logger
}

func NewHandler(logger *zap.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}

var notes []string

func (h *Handler) CreateString(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug(fmt.Sprint("Request time:", time.Now()))

	initTime := time.Now()
	var code int

	defer func() {
		since := time.Since(initTime)

		h.logger.Info(fmt.Sprint("Status code:", code))
		h.logger.Debug(fmt.Sprint("Response time:", time.Now()))
		h.logger.Debug(fmt.Sprint("Handler execution time:", since))
	}()

	h.logger.Info("'CreateString' endpoint with 'POST' method")

	var note string
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		code = http.StatusBadRequest
		http.Error(w, err.Error(), code)
		h.logger.Error(err.Error())
		return
	}

	if note == "" {
		err := errors.New("field is empty")
		code = http.StatusBadRequest
		http.Error(w, err.Error(), code)
		h.logger.Error(err.Error())
		return
	}

	notes = append(notes, note)
	h.logger.Info(fmt.Sprint("Slice length:", len(notes)))

	code = http.StatusCreated
	w.WriteHeader(code)
	if _, err := w.Write([]byte(note)); err != nil {
		h.logger.Error(err.Error())
		return
	}
}

func (h *Handler) GetStrings(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug(fmt.Sprint("Request time:", time.Now()))

	initTime := time.Now()
	var code int

	defer func() {
		since := time.Since(initTime)

		h.logger.Info(fmt.Sprint("Status code:", code))
		h.logger.Debug(fmt.Sprint("Response time:", time.Now()))
		h.logger.Debug(fmt.Sprint("Handler execution time:", since))
	}()

	h.logger.Info("'GetString' endpoint with 'GET' method")

	b, err := json.MarshalIndent(notes, "", "    ")
	if err != nil {
		h.logger.Error(err.Error())
		panic(err)
	}

	code = http.StatusOK
	w.WriteHeader(code)
	if _, err := w.Write(b); err != nil {
		h.logger.Error(err.Error())
		return
	}
}

func main() {
	// logger
	logger, logFileClose, err := logging.NewLogger("info")
	if err != nil {
		panic(err)
	}
	defer logFileClose()

	// handlers
	handlers := NewHandler(logger)

	// routing
	router := mux.NewRouter()

	router.HandleFunc("/strings", handlers.CreateString).Methods("POST")
	router.HandleFunc("/strings", handlers.GetStrings).Methods("GET")

	err = http.ListenAndServe(":9091", router)

	if errors.Is(err, http.ErrServerClosed) {
		return
	} else {
		panic(err)
	}
}
