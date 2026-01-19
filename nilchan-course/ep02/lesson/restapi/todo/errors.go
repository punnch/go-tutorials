package todo

import "errors"

var ErrTaskNotFound error = errors.New("task not found")
var ErrTaskAlreadyExists error = errors.New("task already exists")
