package todo

// main goal: represent all operations with task fields

import "time"

type Task struct {
	Description string
	Text        string
	Completed   bool

	CreatedAt   time.Time
	CompletedAt *time.Time
}

func NewTask(description string, text string) Task {
	return Task{
		Description: description,
		Text:        text,
		Completed:   false,

		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
}

func (t *Task) Complete() {
	completedTime := time.Now()

	t.Completed = true
	t.CompletedAt = &completedTime
}
