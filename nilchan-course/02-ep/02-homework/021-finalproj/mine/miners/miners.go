package miners

import (
	"coalcompany/mine"
	"context"
	"sync"
	"time"
)

/*
- Маленький шахтёр:
  - Оплата труда: 5 угля
  - Энергии хватит на: 30 добыч угля
  - Получает за одну добычу: 1 уголь
  - Перерыв между добычами: 3 секунды
  - Характеристики с каждой добычей не прогрессируют
*/
type MinerInfo struct {
	Cost    mine.Coal
	Energy  int
	GetCoal mine.Coal

	Interval time.Duration
	Progress mine.Coal
}

type Miner struct {
	info MinerInfo

	mtx sync.Mutex
}

func (m *Miner) Run(ctx context.Context) <-chan mine.Coal

func (m *Miner) Info() MinerInfo
