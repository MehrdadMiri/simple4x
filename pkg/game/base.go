package game

import (
	"time"

	"tidy/pkg/debugger"
)

type Game struct {
	Rows, Columns int
	TickerChan    <-chan time.Time
	Debugger      *debugger.Debugger
}

func NewGame(rows, columns int) *Game {
	g := &Game{
		Rows:       rows,
		Columns:    columns,
		TickerChan: time.NewTicker(100 * time.Millisecond).C,
		Debugger:   debugger.NewDebugger(),
	}
	g.Debugger.Set("clock", 0)
	return g
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
