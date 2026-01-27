package enterprise

import (
	"coalcompany/mine"
	"coalcompany/mine/miners"

	"context"
)

// main goal: to interact with whole company (add, get funcs)

// Coal — уголь
// MinerInfo — информация о шахтёре: количество энергии, класс, и т.д.
type Miner interface {
	Run(ctx context.Context) <-chan mine.Coal
	Info() miners.MinerInfo
}
