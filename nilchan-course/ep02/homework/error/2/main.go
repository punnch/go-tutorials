package main

import (
	"fmt"
	"study/bank"
)

func main() {
	// Create a user
	user := bank.NewUser()

	// Add user fields (encapsulation)
	user.AddName("Dima")
	user.AddBalance(100)

	// Infinity loop before any error encountered
	for {
		// ShowBalance()
		usd, err := user.ShowBalance()
		if err != nil {
			fmt.Println("ShowBalance() error!\nDescription:", err)
			break
		} else {
			fmt.Println("Showbalance() was successful\nBalance:", usd)
		}

		fmt.Println()

		// WithdrawMoney()
		err = user.WithdrawMoney(20)
		if err != nil {
			fmt.Println("WithdrawMoney() error!\nDescription:", err)
			break
		} else {
			fmt.Println("WithdrawMoney() was successful")
		}

		fmt.Println()

		// ShowBalance()
		usd, err = user.ShowBalance()
		if err != nil {
			fmt.Println("ShowBalance() error!\nDescription:", err)
			break
		} else {
			fmt.Println("Showbalance() was successful\nBalance:", usd)
		}

		fmt.Println()

		// Pay()
		err = user.Pay(200)
		if err != nil {
			fmt.Println("Pay() error!\nDescription:", err)
			break
		} else {
			fmt.Println("Pay() was successful")
		}

		fmt.Println()

		// ShowBalance()
		usd, err = user.ShowBalance()
		if err != nil {
			fmt.Println("ShowBalance() error!\nDescription:", err)
			break
		} else {
			fmt.Println("Showbalance() was successful\nBalance:", usd)
		}
	}
}
