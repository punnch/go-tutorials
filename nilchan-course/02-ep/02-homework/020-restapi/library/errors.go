package library

import "errors"

var ErrBookAlreadyExist error = errors.New("book already exist")
var ErrBookNotFound error = errors.New("book not found")
