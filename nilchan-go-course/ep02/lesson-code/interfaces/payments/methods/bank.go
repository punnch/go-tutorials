package methods

import (
	"fmt"
	"math/rand"
)

type Bank struct{}

func NewBank() Bank {
	return Bank{}
}

func (b Bank) Pay(usd int) int {
	fmt.Println("Оплата через банк")
	fmt.Println("Размер оплаты:", usd, "долларов")

	return rand.Int()
}

func (b Bank) Cancel(id int) {
	fmt.Println("Отмена по ID", id)
}
