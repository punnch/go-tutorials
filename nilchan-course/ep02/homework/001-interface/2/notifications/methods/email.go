package methods

import "fmt"

type Email struct{}

// Constructor
func NewEmail() Email {
	return Email{}
}

// Accepts Notifier interface
func (e Email) Send(to, message string) {
	fmt.Println("Отправка через email!")
	fmt.Println("Получатель:", to)
	fmt.Println("Сообщение:", message)
	fmt.Println("-------------------")
}
