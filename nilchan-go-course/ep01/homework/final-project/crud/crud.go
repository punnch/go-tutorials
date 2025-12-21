package crud

import (
	"fmt"
	"time"

	"todo/db"

	"github.com/k0kubun/pp"
)

// Events
func AddEvent(inpText string, errText string) {
	event := db.Event{
		Text: inpText,
		Desc: errText,
		Time: time.Now(),
	}

	db.Events = append(db.Events, event)
}

func EventList() {
	for i, event := range db.Events {
		pp.Println(i+1, event)
	}
}

// Tasks
func AddTask(title string, desc string) {
	task := db.Task{
		ID:    db.IdCounter,
		Title: title,
		Desc:  desc,
		Time:  time.Now(),
		Done:  false,
	}

	db.Tasks = append(db.Tasks, task)
	db.IdCounter++
}

func TaskList() {
	for _, task := range db.Tasks {
		status := "❌"

		// if completed (true)
		if task.Done {
			status = "✅"
		}

		fmt.Printf(
			"%d. %s | %s %s [%s]\n",
			task.ID,
			task.Title,
			task.Desc,
			task.Time.Format("15:04:05 01-02-2006"),
			status,
		)
	}
}

func DeleteTasks(title string) string {
	for i, task := range db.Tasks {
		if task.Title == title {
			db.Tasks = append(db.Tasks[:i], db.Tasks[i+1:]...)
			return ""
		}
	}
	return "Задача с таким заголовком не найдена!"
}

func ChangeStatus(title string) string {
	for _, task := range db.Tasks {
		if task.Title == title {
			// Ставим противоположный статус
			task.Done = !task.Done
			return ""
		}
	}
	return "Задача с таким заголовком не найдена!"
}

func Help() {
	fmt.Println("- help — позволяет узнать доступные команды и их формат")
	fmt.Println("- add {заголовок задачи 1 слово} {текст задачи 1-неск. слов} — позволяет добавлять задачи в список задач")
	fmt.Println("- list — позволяет получить полный список всех задач")
	fmt.Println("- del {заголовок существующей задачи} — позволяет удалить задачу по её заголовку")
	fmt.Println("- done {заголовок существующей задачи} — позволяет отметить задачу как выполненную")
	fmt.Println("- events — позволяет получить список всех событий")
	fmt.Println("- exit — позволяет завершить выполнение программы")
}
