package coordinates

import (
	"fmt"
	"math/rand"
)

func RandomCoordinates() string {
	max := 2000
	min := -2000

	x := rand.Intn(max-min+1) + min
	y := rand.Intn(max-min) + min
	z := rand.Intn(max-min) + min

	return fmt.Sprint(x, y, z)
}
