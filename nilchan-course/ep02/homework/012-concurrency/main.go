package main

import (
	"concurrency/geo"
	"concurrency/sensors"
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	airPressureSlice := []int{}
	airMoisureSlice := []int{}
	seismicActivitySlice := []int{}
	mtx := sync.Mutex{}

	// Set sensor count
	airPressureCount := 5
	airMoisureCount := 5
	seismicActivityCount := 5

	// Create Contexts
	pressureCtx, pressureCancel := context.WithCancel(context.Background())
	moisureCtx, moisureCancel := context.WithCancel(context.Background())
	activityCtx, activityCancel := context.WithCancel(context.Background())

	// Create Channels
	pressureCh := sensors.PressurePool(pressureCtx, airPressureCount)
	moisureCh := sensors.AirPool(moisureCtx, airMoisureCount)
	activityCh := sensors.ActivityPool(activityCtx, seismicActivityCount)

	// Close contexts after certain time
	go func() {
		time.Sleep(3 * time.Second)
		pressureCancel()
		fmt.Println("ДАТЧИКИ ДАВЛЕНИЯ ВОЗДУХА ОТКЛЮЧАЮТСЯ!!!")
	}()

	go func() {
		time.Sleep(5 * time.Second)
		moisureCancel()
		fmt.Println("ДАТЧИКИ ВЛАЖНОСТИ ВОЗДУХА ОТКЛЮЧАЮТСЯ!!!")
	}()

	go func() {
		time.Sleep(7 * time.Second)
		activityCancel()
		fmt.Println("ДАТЧИКИ СЕЙСМИЧЕСКОЙ АКТИВНОСТИ ОТКЛЮЧАЮТСЯ!!!")
	}()

	wg := &sync.WaitGroup{}

	initTime := time.Now()

	// Get values from channels
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range pressureCh {
			mtx.Lock()
			airPressureSlice = append(airPressureSlice, v)
			mtx.Unlock()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range moisureCh {
			mtx.Lock()
			airMoisureSlice = append(airMoisureSlice, v)
			mtx.Unlock()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range activityCh {
			mtx.Lock()
			seismicActivitySlice = append(seismicActivitySlice, v)
			mtx.Unlock()
		}
	}()

	wg.Wait()

	// Get Average
	mtx.Lock()
	airPressureAverage := geo.Average(airPressureSlice)
	mtx.Unlock()

	mtx.Lock()
	airMoisureAverage := geo.Average(airMoisureSlice)
	mtx.Unlock()

	mtx.Lock()
	seismicActivityAverage := geo.Average(seismicActivitySlice)
	mtx.Unlock()

	// Output
	fmt.Println("---------------------------------------------")
	fmt.Println("Время выполнения:", time.Since(initTime))
	fmt.Println("Среднее арифмитическое давления:", airPressureAverage, "мм рт. ст.")
	fmt.Println("Среднее арифмитическое влажности воздуха:", airMoisureAverage, "%")
	fmt.Println("Среднее арифмитическое сейсмической активности:", seismicActivityAverage, "баллов")
}
