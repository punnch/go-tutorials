package figures

import "fmt"

type shape interface {
	Area() float64
	Perim() float64
}

func PrintShape(s shape) {
	fmt.Println("Area:", s.Area())
	fmt.Println("Perim:", s.Perim())
	fmt.Println("---")
}
