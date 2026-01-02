package main

import (
	"errors"
	"fmt"

	"github.com/k0kubun/pp"
)

type User struct {
	Name    string
	Balance int
}

func Pay(u *User, usd int) error {
	if u.Balance-usd < 0 {
		return errors.New("insufficient money")
	}

	u.Balance -= usd

	return nil
}

func main() {
	user := User{
		Name:    "Dima",
		Balance: 200,
	}

	pp.Println("Balance before a payment:", user.Balance)
	err := Pay(&user, 150)
	pp.Println("After:", user.Balance)

	if err != nil {
		fmt.Println("Payment error! Reason:", err)
	} else {
		fmt.Println("Payment was successful!")
	}
}
