package main

import (
	"fmt"
	"log"

	"github.com/punnch/calendar"
)

func main() {
	event := calendar.Event{}

	// Year
	err := event.SetYear(2026)
	if err != nil {
		log.Fatal(err)
	}

	// Month
	err = event.SetMonth(1)
	if err != nil {
		log.Fatal(err)
	}

	// Day
	err = event.SetDay(1)
	if err != nil {
		log.Fatal(err)
	}

	// Title
	err = event.SetTitle("New year!!!")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(event)
	fmt.Println("Year:", event.Year())
	fmt.Println("Month:", event.Month())
	fmt.Println("Day:", event.Day())
	fmt.Println("Title:", event.Title())
}
