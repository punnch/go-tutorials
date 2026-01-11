package geo

func Average(arr []int) int {
	var sum int
	length := len(arr)

	for _, v := range arr {
		sum += v
	}

	average := sum / length

	return average
}
