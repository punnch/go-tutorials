package mine

import "time"

type Coal int

type Statistics struct {
	CoalAmount Coal
	TotalTime  time.Duration
}
