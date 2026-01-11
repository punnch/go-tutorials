package sensors

import (
	"concurrency/geo"
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func air(
	ctx context.Context,
	ch chan<- int,
	wg *sync.WaitGroup,
	n int,
	moisure int,
	coordinates string,
) {
	defer wg.Done()

	for {
		fmt.Println("Я датчик влажности воздуха", n, "Анализирую...")

		select {
		case <-ctx.Done():
			fmt.Println("Я датчик влажности воздуха", n, "Закончил работу! (отключился)")
			return
		case <-time.After(time.Second):
			fmt.Println("Я датчик влажности воздуха", n, "Измерил:", moisure, "%")
		}

		select {
		case <-ctx.Done():
			fmt.Println("Я датчик влажности воздуха", n, "Закончил работу! (отключился)")
			return
		case ch <- moisure:
			fmt.Println("Я датчик влажности воздуха", n, "Закончил! Координаты:", coordinates, "!")
		}
	}
}

func AirPool(ctx context.Context, sensorCount int) <-chan int {
	moisureCh := make(chan int)
	wg := &sync.WaitGroup{}

	for i := 1; i <= sensorCount; i++ {
		wg.Add(1)
		go air(ctx, moisureCh, wg, i, getAirMoisure(), geo.RandomCoordinates())
	}

	go func() {
		wg.Wait()
		close(moisureCh)
	}()

	return moisureCh
}

func getAirMoisure() int {
	min := 20
	max := 80

	randomAirMoisure := rand.IntN(max-min+1) + min
	return randomAirMoisure
}
