package air

import (
	"concurrency/coordinates"
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
)

func air(
	ctx context.Context,
	ch chan<- int,
	wg *sync.WaitGroup,
	n int,
	airMoisure int,
	coordinates string,
) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Я датчик влажности воздуха", n, "Закончил работу! (отключился)")
			return
		default:
			fmt.Println("Я датчик влажности воздуха", n, "Анализирую...")
			fmt.Println("Я датчик влажности воздуха", n, "Измерил:", airMoisure, "%")

			ch <- airMoisure
			fmt.Println("Я датчик влажности воздуха", n, "Закончил! Координаты:", coordinates, "!")
		}
	}
}

func AirPool(ctx context.Context, sensorCount int) <-chan int {
	airCh := make(chan int)
	wg := &sync.WaitGroup{}

	for i := 1; i <= sensorCount; i++ {
		wg.Add(1)
		go air(ctx, airCh, wg, i, getAirMoisure(), coordinates.RandomCoordinates())
	}

	go func() {
		wg.Wait()
		close(airCh)
	}()

	return airCh
}

func getAirMoisure() int {
	min := 20
	max := 80

	randomAirMoisure := rand.IntN(max-min+1) + min
	return randomAirMoisure
}
