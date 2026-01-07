package main

import (
	"fmt"
	"packages_and_modules/user"

	"github.com/k0kubun/pp"
)

func main() {
	u := user.User{}

	user1 := u.NewUser("Dima", 15)
	fmt.Println("User1:", user1)

	u.ChangeAge(16)
	fmt.Println(u.GetAge())

	u.ChangeName("Patrik")

	fmt.Println("User1:", user1)

	fmt.Println(u)

	pp.Println(user1)
}
