package activity

import (
	"concurrency/coordinates"
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
)

func activity(
	ctx context.Context,
	wg *sync.WaitGroup,
	ch chan<- int,
	n int,
	activity int,
	coordinates string,
) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Я датчик сейсмической активности", n, "Закончил работу! (отключился)")
		default:
			fmt.Println("Я датчик сейсмической активности", n, "Анализирую...")
			fmt.Println("Я датчик сейсмической активности", n, "Измерил:", activity)

			ch <- activity
			fmt.Println("Я датчик сейсмической активности", n, "Законил! Координаты:", coordinates, "!")
		}
	}
}

func ActivityPool(ctx context.Context, sensorCount int) <-chan int {
	activityCh := make(chan int)
	wg := &sync.WaitGroup{}

	for i := 1; i <= sensorCount; i++ {
		wg.Add(1)
		go activity(ctx, wg, activityCh, sensorCount, getSeismicActivity(), coordinates.RandomCoordinates())
	}

	go func() {
		wg.Wait()
		close(activityCh)
	}()

	return activityCh
}

func getSeismicActivity() int {
	max := 12

	randomSeismicActivity := rand.IntN(max + 1) // 0 - 12

	return randomSeismicActivity
}
