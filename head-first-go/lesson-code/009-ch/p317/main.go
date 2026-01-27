package main

import "fmt"

type rub float64
type uah float64
type usd float64

// 1 for rub
func (r rub) toUsd() usd {
	return usd(r / 81.38)
}
func (r rub) toUah() uah {
	return uah(r / 1.95)
}

// 2 for usd
func (u usd) toRub() rub {
	return rub(u * 81.38)
}
func (u usd) toUah() uah {
	return uah(u * 41.67)
}

// 3 for uah
func (u uah) toUsd() usd {
	return usd(u / 41.67)
}
func (u uah) toRub() uah {
	return uah(u * 1.95)
}

func main() {
	rubles := rub(250)
	usd := usd(14.39)
	uah := uah(144.4)

	fmt.Printf("%.2f rub in usd is %.2f\n", rubles, rubles.toUsd())
	fmt.Printf("%.2f usd in rub is %.2f\n", usd, usd.toRub())
	fmt.Printf("%.2f uah in rub is %.2f\n", uah, uah.toRub())
}
