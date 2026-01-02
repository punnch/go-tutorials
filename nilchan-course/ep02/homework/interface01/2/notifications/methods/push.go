package methods

import "fmt"

type Push struct{}

// Constructor
func NewPush() Push {
	return Push{}
}

// Accepts Notifier interface
func (p Push) Send(to, message string) {
	fmt.Println("Отправка через Push!")
	fmt.Println("Получатель:", to)
	fmt.Println("Сообщение:", message)
	fmt.Println("-------------------")
}
