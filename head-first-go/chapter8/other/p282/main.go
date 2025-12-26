package main

import (
	"fmt"

	"github.com/punnch/database"
)

func main() {
	sub1 := database.Employee{}
	sub1.Name = "Cumpot"
	sub1.Salary = 10000
	sub1.Country = "Ukraine"
	sub1.City = "Kirovograd"

	// 1 Employee
	fmt.Println("*Employee*")
	fmt.Println("Name:", sub1.Name)
	fmt.Printf("Salary: %.f$\n", sub1.Salary)
	fmt.Println("Country:", sub1.Country)
	fmt.Println("City:", sub1.City)
}
