package db

import (
	"time"
)

// Tasks
type Task struct {
	ID       int
	Title    string
	Desc     string
	Time     time.Time
	Done     bool
	DoneTime time.Time
}

var Tasks []Task = []Task{}
var IdCounter int = 0

// Events
type Event struct {
	Text string
	Desc string
	Time time.Time
}

var Events []Event
