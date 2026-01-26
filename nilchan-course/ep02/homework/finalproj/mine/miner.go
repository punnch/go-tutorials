package mine

import "time"

type MinerInfo struct {
	Class string

	Cost   int
	Energy int
	Coal   string

	Break        time.Time
	CoalProgress int
}
