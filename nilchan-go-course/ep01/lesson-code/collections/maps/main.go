package main

import "fmt"

func main() {
	/* shop := map[string]float64{
		"bread":     0.99,
		"juice":     1,
		"ice_cream": 3,
		"chocolate": 5,
	}

	fmt.Println(shop)
	fmt.Println("cost:", shop["chocolate"], "$")

	shop["cake"] = 10

	fmt.Println(shop)
	fmt.Println("cost:", shop["cake"], "$")

	// Вызываю не созданный ключ
	fmt.Println(shop["not_created"])

	// Проверяю вывод не созданного ключа
	burg, ok := shop["burger"]
	fmt.Println(burg, ok) /* if ok is false it means variable isn't created

	fmt.Println("--------")
	fmt.Println("До:")
	fmt.Println(shop)

	for k, _ := range shop {
		shop[k]++
	}

	fmt.Println("--------")
	fmt.Println("После:")
	fmt.Println(shop) */

	weather := make(map[int]int)

	fmt.Println("---------")
	fmt.Println("До:")
	fmt.Println(weather)
	fmt.Println("")

	weather[1] = 25
	weather[2] = 35
	weather[3] = 28

	fmt.Println("---------")
	fmt.Println("После:")
	fmt.Println(weather)
	fmt.Println("")
}
