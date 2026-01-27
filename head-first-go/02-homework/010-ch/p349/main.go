package main

import (
	"fmt"
	"log"

	"github.com/punnch/geo"
)

func main() {
	location := geo.Landmark{}

	err := location.SetName("The googleplex")
	if err != nil {
		log.Fatal(err)
	}

	err = location.SetLatitude(-37.31)
	if err != nil {
		log.Fatal(err)
	}

	err = location.SetLongitude(150.61)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(location.Name())
	fmt.Println(location.Latitude())
	fmt.Println(location.Longitude())
}
