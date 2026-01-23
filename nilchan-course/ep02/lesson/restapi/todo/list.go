package todo

// main purpose: represent an opportunity to ineract with tasks

import (
	"maps"
	"sync"
)

type List struct {
	tasks map[string]Task
	mtx   sync.RWMutex
}

func NewList() *List {
	return &List{
		tasks: make(map[string]Task),
	}
}

func (l *List) AddTask(task Task) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	if _, ok := l.tasks[task.Title]; ok {
		return ErrTaskAlreadyExists
	}

	l.tasks[task.Title] = task

	return nil
}

func (l *List) GetTask(title string) (Task, error) {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	task, ok := l.tasks[title]
	if !ok {
		return Task{}, ErrTaskNotFound
	}

	return task, nil
}

func (l *List) GetAllTasks() map[string]Task {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	tmp := make(map[string]Task, len(l.tasks))

	maps.Copy(tmp, l.tasks)

	return tmp
}

func (l *List) GetAllUncompletedTasks() map[string]Task {
	uncompletedTasks := make(map[string]Task)

	l.mtx.RLock()
	defer l.mtx.RUnlock()

	for title, task := range l.tasks {
		if !task.Completed {
			uncompletedTasks[title] = task
		}
	}

	return uncompletedTasks
}

func (l *List) CompleteTask(title string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	task, ok := l.tasks[title]
	if !ok {
		return ErrTaskNotFound
	}

	task.Complete()

	l.tasks[title] = task

	return nil
}

func (l *List) UncompleteTask(title string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	task, ok := l.tasks[title]
	if !ok {
		return ErrTaskNotFound
	}

	task.Uncomplete()

	l.tasks[title] = task

	return nil
}

func (l *List) DeleteTask(title string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	_, ok := l.tasks[title]
	if !ok {
		return ErrTaskNotFound
	}

	delete(l.tasks, title)

	return nil
}
