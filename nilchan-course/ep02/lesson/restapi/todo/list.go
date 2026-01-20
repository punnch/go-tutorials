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

func (l *List) AddTask(task Task) error {
	// check if task already exists
	if _, ok := l.tasks[task.Title]; ok {
		return ErrTaskAlreadyExists
	}

	l.tasks[task.Title] = task

	return nil
}

func (l *List) ListAllTasks() map[string]Task {
	// return a map copy
	tmp := make(map[string]Task, len(l.tasks))

	maps.Copy(tmp, l.tasks)

	return tmp
}

func (l *List) ListAllNotCompletedTasks() map[string]Task {
	notCompletedTasks := make(map[string]Task)

	for title, task := range l.tasks {
		if !task.Completed {
			notCompletedTasks[title] = task
		}
	}

	return notCompletedTasks
}

func (l *List) CompleteTask(title string) error {
	task, ok := l.tasks[title]
	if !ok {
		return ErrTaskNotFound
	}

	task.Complete()

	l.tasks[title] = task

	return nil
}

func (l *List) DeleteTask(title string) error {
	_, ok := l.tasks[title]
	if !ok {
		return ErrTaskNotFound
	}

	delete(l.tasks, title)

	return nil
}
