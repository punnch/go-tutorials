package todo

// main purpose: represent an opportunity to ineract with tasks

import "maps"

type List struct {
	tasks map[string]Task
}

func NewList() *List {
	return &List{
		tasks: make(map[string]Task),
	}
}

// check if task already exists
func (l *List) AddTask(task Task) error {
	if _, ok := l.tasks[task.Description]; ok {
		return ErrTaskAlreadyExists
	}

	l.tasks[task.Description] = task

	return nil
}

// return a map copy
func (l *List) ListTasks() map[string]Task {
	tmp := make(map[string]Task, len(l.tasks))

	maps.Copy(tmp, l.tasks)

	return tmp
}

func (l *List) CompleteTask(description string) error {
	task, ok := l.tasks[description]
	if !ok {
		return ErrTaskNotFound
	}

	task.Complete()

	l.tasks[description] = task

	return nil
}

func (l *List) DeleteTask(description string) error {
	_, ok := l.tasks[description]
	if !ok {
		return ErrTaskNotFound
	}

	delete(l.tasks, description)

	return nil
}
