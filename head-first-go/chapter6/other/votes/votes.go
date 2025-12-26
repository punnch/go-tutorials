package main

import (
	"fmt"
	"log"

	"github.com/punnch/datafile"
)

func main() {
	var names []string
	var counts []int

	lines, err := datafile.GetStrings("../datafile/votes.txt")
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range lines {
		matched := false

		for i, name := range names {
			if name == line {
				counts[i]++
				matched = true
			}
		}

		if !matched {
			names = append(names, line)
			counts = append(counts, 1)
		}
	}

	for i, name := range names {
		fmt.Printf("%s: %d\n", name, counts[i])
	}
}
