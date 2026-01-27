package main

import "fmt"

type user struct {
	name   string
	rate   float64
	rating float64
	active bool
}

func printInfo(s *user) {
	fmt.Println("Name:", s.name)
	fmt.Printf("Monthly rate cost: %.f$\n", s.rate)
	fmt.Println("Rating:", s.rating)
	fmt.Println("Subscription:", s.active)
}

func defaultUser(name string) *user {
	u := user{}
	u.name = name
	u.rate = 99
	u.rating = 5
	u.active = true

	return &u
}

func applyDiscount(s *user) {
	s.rate = 49
}

func main() {
	sub1 := defaultUser("cumpot")
	sub1.rating = 0
	printInfo(sub1)

	fmt.Println("") // indent

	sub2 := defaultUser("Punnch")
	applyDiscount(sub2)
	printInfo(sub2)
}
