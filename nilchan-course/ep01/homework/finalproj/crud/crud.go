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
	return "No task found with the provided title!"
}

func ChangeStatus(title string) string {
	for _, task := range db.Tasks {
		if task.Title == title {
			// toggle
			task.Done = !task.Done
			return ""
		}
	}
	return "No task found with the provided title!"
}

func Help() {
	fmt.Println("- help — show commands")
	fmt.Println("- add {title (1 word)} {description} — make a note")
	fmt.Println("- list — get uncompleted notes")
	fmt.Println("- del {title} — delete a note by title")
	fmt.Println("- done {title} — switch status to 'done'")
	fmt.Println("- events — get all events")
	fmt.Println("- exit — close the sessions")
}
