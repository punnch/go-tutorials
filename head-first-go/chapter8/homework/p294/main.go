package main

import (
	"fmt"
	"googleplex/geo"
)

func main() {
	location := geo.Landmark{}
	location.Name = "The Googleplex"
	location.Latitude = 31.88
	location.Longitude = -100.5

	fmt.Println(location)
}
