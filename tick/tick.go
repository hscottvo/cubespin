package tick

import (
	"fmt"
	"time"
)

type Ticker struct {
	fps     int
	Runtime int
	ticker  *time.Ticker
}

func NewTicker(fps int, runtime int) *Ticker {
	// gets in seconds
	period := time.Second / time.Duration(fps)
	fmt.Println(period)
	ticker := time.NewTicker(period)

	t := Ticker{fps, runtime, ticker}

	return &t
}

func (t Ticker) Tick() <-chan time.Time {
	return t.ticker.C
}
