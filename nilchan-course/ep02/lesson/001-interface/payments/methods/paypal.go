package methods

import (
	"fmt"
	"math/rand"
)

type PayPal struct{}

func NewPayPal() PayPal {
	return PayPal{}
}

func (p PayPal) Pay(usd int) int {
	fmt.Println("Оплата через PayPal")
	fmt.Println("Размер оплаты:", usd, "usd")

	return rand.Int()
}

func (p PayPal) Cancel(id int) {
	fmt.Println("Отмена крипто-операции по ID", id)
}
