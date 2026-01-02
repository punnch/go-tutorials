// package bank is to make easier bank operations:

// - withdraw money
// - show balance
// - pay

// It has these features:

// 1. Operation error has 30 % chance
// 2. Error handling of each function (except NewUser())

package bank

import (
	"errors"
	"math/rand"
)

// NewUser() constructor
func NewUser() User {
	return User{}
}

// AddName() add a name field
func (u *User) AddName(name string) error {
	if name == "" {
		return errors.New("empty name field")
	}

	u.name = name

	return nil
}

// AddName() add a balance field
func (u *User) AddBalance(usd int) error {
	if usd < 0 {
		return errors.New("negative balance")
	}

	u.balance = usd

	return nil
}

// randomError() do 30% error chance
func randomError() error {
	num := rand.Intn(10)
	if num < 3 {
		return errors.New("operation error, try again later")
	}

	return nil
}

// WithdrawMoney() withdraw money from user's balance
func (u *User) WithdrawMoney(usd int) error {
	if err := randomError(); err != nil {
		return err
	}

	if u.balance-usd < 0 {
		return errors.New("not enough money to withdraw, try to get less money")
	}

	u.balance -= usd

	return nil
}

// ShowBalance() returns current user balance
func (u *User) ShowBalance() (int, error) {
	if err := randomError(); err != nil {
		return 0, err
	}

	return u.balance, nil
}

// Pay() make a payment from user's balance
func (u *User) Pay(usd int) error {
	if err := randomError(); err != nil {
		return err
	}

	if u.balance-usd < 0 {
		return errors.New("not enough money to pay")
	}

	u.balance -= usd

	return nil
}
