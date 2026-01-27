package methods

import "fmt"

type Sms struct{}

// Constructor
func NewSms() Sms {
	return Sms{}
}

// Accepts Notifier interface
func (s Sms) Send(to, message string) {
	fmt.Println("Отправка через SMS!")
	fmt.Println("Получатель:", to)
	fmt.Println("Сообщение:", message)
	fmt.Println("-------------------")
}
