package pressure

import (
	"concurrency/coordinates"
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
)

func pressure(
	ctx context.Context,
	ch chan<- int,
	wg *sync.WaitGroup,
	n int,
	pressure int,
	cordinates string,
) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Я датчик давления воздуха", n, "Закончил работу! (отключился)")
			return
		default:
			fmt.Println("Я датчик давления воздуха", n, "Анализирую...")
			fmt.Println("Я датчик давления воздуха", n, "Измерил:", pressure)

			ch <- pressure
			fmt.Println("Я датчик давления воздуха", n, "Закончил! Координаты:", cordinates, "!")
		}
	}
}

func PressurePool(ctx context.Context, sensorCount int) <-chan int {
	pressureCh := make(chan int)
	wg := &sync.WaitGroup{}

	for i := 1; i <= sensorCount; i++ {
		wg.Add(1)
		go pressure(ctx, pressureCh, wg, i, getPressure(), coordinates.RandomCoordinates())
	}

	go func() {
		wg.Wait()
		close(pressureCh)
	}()

	return pressureCh
}

func getPressure() int {
	max := 780
	min := 720

	randomPressure := rand.IntN(max-min+1) + min
	return randomPressure
}
