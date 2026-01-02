package main

import (
	"fmt"
	"study/payments"
	"study/payments/methods"

	"github.com/k0kubun/pp"
)

func main() {
	method := methods.NewCrypto()
	paymentModule := payments.NewPaymentModule(method)

	id := paymentModule.Pay("burger", 10)
	paymentModule.Pay("coke", 2)

	info := paymentModule.Info(id)
	fmt.Println(info)

	paymentModule.Cancel(id)

	info = paymentModule.Info(id)
	fmt.Println(info)

	pp.Println(paymentModule.AllInfo())
}
