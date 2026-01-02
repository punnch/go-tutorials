package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"todo/crud"
)

func main() {
	var eventDesc string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Добро пожаловать в ToDo!")
	fmt.Print("(Чтобы получить список доступных команд, введите 'help')\n\n")

	for {
		fmt.Print("Введите команду: ")

		// Checking input ability
		if ok := scanner.Scan(); !ok {
			log.Fatal(scanner.Err())
		}

		text := scanner.Text()

		fields := strings.Fields(text)
		if len(fields) == 0 {
			eventDesc = "Пустой запрос!"
			fmt.Println(eventDesc)
			crud.AddEvent("", eventDesc)
			continue
		}

		cmd := fields[0]

		var desc string
		if len(fields) >= 3 {
			desc = strings.Join(fields[2:], " ")
		}

		if cmd == "exit" {
			if len(fields) != 1 {
				eventDesc = "Неверный формат!"
				fmt.Println(eventDesc)
				crud.AddEvent(text, eventDesc)
				continue
			}
			fmt.Println("Вы закрыли приложение")
			break
		}

		if cmd == "add" {
			if len(fields) < 3 {
				eventDesc = "Неверный формат!"
				fmt.Println(eventDesc)
				crud.AddEvent(text, eventDesc)
				continue
			}
			crud.AddTask(fields[1], desc)
			eventDesc = ""
			fmt.Println(eventDesc)
			crud.AddEvent(text, eventDesc)
			continue
		}

		if cmd == "del" {
			if len(fields) != 2 {
				eventDesc = "Неверный формат!"
				fmt.Println(eventDesc)
				crud.AddEvent(text, eventDesc)
				continue
			}
			eventDesc = crud.DeleteTasks(fields[1])
			if eventDesc != "" {
				fmt.Println(eventDesc)
				crud.AddEvent(text, eventDesc)
				continue
			}
			eventDesc = ""
			fmt.Println(eventDesc)
			crud.AddEvent(text, eventDesc)
			continue
		}

		if cmd == "done" {
			if len(fields) != 2 {
				eventDesc = "Неверный формат!"
				fmt.Println(eventDesc)
				crud.AddEvent(text, eventDesc)
				continue
			}
			crud.ChangeStatus(fields[1])
			if eventDesc != "" {
				fmt.Println(eventDesc)
				crud.AddEvent(text, eventDesc)
				continue
			}
			eventDesc = ""
			fmt.Println(eventDesc)
			crud.AddEvent(text, eventDesc)
			continue
		}

		if cmd == "list" {
			if len(fields) != 1 {
				eventDesc = "Неверный формат!"
				fmt.Println(eventDesc)
				crud.AddEvent(text, eventDesc)
				continue
			}
			crud.TaskList()
			eventDesc = ""
			fmt.Println(eventDesc)
			crud.AddEvent(text, eventDesc)
			continue
		}

		if cmd == "help" {
			if len(fields) != 1 {
				eventDesc = "Неверный формат!"
				fmt.Println(eventDesc)
				crud.AddEvent(text, eventDesc)
				continue
			}
			crud.Help()
			eventDesc = ""
			fmt.Println(eventDesc)
			crud.AddEvent(text, eventDesc)
			continue
		}

		if cmd == "events" {
			if len(fields) != 1 {
				eventDesc = "Неверный формат!"
				fmt.Println(eventDesc)
				crud.AddEvent(text, eventDesc)
				continue
			}
			crud.EventList()
			eventDesc = ""
			fmt.Println(eventDesc)
			crud.AddEvent(text, eventDesc)
			continue
		}

		eventDesc = "Передана неизвестная команда!"
		fmt.Println(eventDesc)
		crud.AddEvent(text, eventDesc)
	}
}
