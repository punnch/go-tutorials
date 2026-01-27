package sensors

import (
	"concurrency/geo"
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
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
		fmt.Println("Я датчик сейсмической активности", n, "Анализирую...")

		select {
		case <-ctx.Done():
			fmt.Println("Я датчик сейсмической активности", n, "Закончил работу! (отключился)")
			return
		case <-time.After(time.Second):
			fmt.Println("Я датчик сейсмической активности", n, "Измерил:", activity, "баллов")
		}

		select {
		case <-ctx.Done():
			fmt.Println("Я датчик сейсмической активности", n, "Закончил работу! (отключился)")
			return
		case ch <- activity:
			fmt.Println("Я датчик сейсмической активности", n, "Законил! Координаты:", coordinates, "!")
		}
	}
}

func ActivityPool(ctx context.Context, sensorCount int) <-chan int {
	activityCh := make(chan int)
	wg := &sync.WaitGroup{}

	for i := 1; i <= sensorCount; i++ {
		wg.Add(1)
		go activity(ctx, wg, activityCh, i, getSeismicActivity(), geo.RandomCoordinates())
	}

	go func() {
		wg.Wait()
		close(activityCh)
	}()

	return activityCh
}

func getSeismicActivity() int {
	max := 12

	randomSeismicActivity := rand.IntN(max + 1)

	return randomSeismicActivity
}
